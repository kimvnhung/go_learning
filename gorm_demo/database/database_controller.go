package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go_learning/gorm_demo/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DatabaseController struct {
	DB           *gorm.DB
	DataFileName string
}

func New() *DatabaseController {

	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		databaseName = "empty_database_name"
	}

	p := &DatabaseController{
		DB:           prepareDatabase(databaseName),
		DataFileName: databaseName,
	}

	return p
}

func prepareDatabase(databaseName string) *gorm.DB {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "1123"
	}
	dsn := fmt.Sprintf("host=%s user=postgres password=1 database=%s port=%s sslmode=disable TimeZone=Asia/Saigon", dbHost, databaseName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.Book{}, &models.Chapter{}, &models.Entry{})

	//init value
	//Main : novel, screenplay, article, speech, poem, research/essay,other

	return db
}

/*
Insert screen problem if has the same nodes
*/
type UpdateType int

const (
	NONE UpdateType = iota
	WITHOUT_ASSOCIATIONS
	ASSOCIATIONS
)

func (dc *DatabaseController) Insert(item any, update UpdateType) error {
	log.Printf("insert %T", item)
	defer log.Printf("insert %T done", item)
	if book, ok := item.(models.Book); ok {
		rs := dc.DB.Create(&book)
		if rs.Error != nil {
			if update == NONE {
				return rs.Error
			}

			log.Println("try to update")

			existed, err := dc.GetBookByUnique(book.Name)
			if err != nil {
				return err
			}

			if update == ASSOCIATIONS {
				// err = dc.DB.Unscoped().Model(&existed).Association("Chapters").Clear()
				// if err != nil {
				// 	return err
				// }
				for _, ch := range book.Chapters {
					rs = dc.DB.Save(&ch)
					if rs.Error != nil {
						return rs.Error
					}
				}
				existed.Chapters = book.Chapters
			}
			existed.Name = book.Name

			log.Printf("existed before save %v", existed)

			rs = dc.DB.Save(&existed)
			if rs.Error != nil {
				return rs.Error
			}
		}
	} else if chapter, ok := item.(models.Chapter); ok {
		rs := dc.DB.Create(&chapter)
		if rs.Error != nil {
			if update == NONE {
				return rs.Error
			}

			log.Println("try to update")

			existed, err := dc.GetChapterByUnique(chapter.BookID, chapter.Index)
			if err != nil {
				return err
			}

			if update == ASSOCIATIONS {
				existed.Entries = chapter.Entries
			}

			existed.Name = chapter.Name

			rs = dc.DB.Save(&existed)
			if rs.Error != nil {
				return rs.Error
			}
		}
	} else if entry, ok := item.(models.Entry); ok {
		rs := dc.DB.Create(&entry)
		if rs.Error != nil {
			if update == NONE {
				return rs.Error
			}

			log.Println("try to update")

			existed, err := dc.GetEntryByUnique(entry.ChapterID, entry.Index)
			if err != nil {
				return err
			}

			existed.Name = entry.Name

			rs = dc.DB.Save(&existed)
			if rs.Error != nil {
				return rs.Error
			}
		}
	}

	return nil
}

/*
	has no option for screen log and screen node log
*/

func (dc *DatabaseController) Prepare(item any) (any, error) {
	log.Printf("prepare %T", item)
	defer log.Printf("prepare %T done", item)

	return nil, errors.New("prepare failed")
}

func (dc *DatabaseController) GetBookByUnique(bookName string) (models.Book, error) {
	var book models.Book

	rs := dc.DB.Preload("Chapters."+clause.Associations).Where("name = ?", bookName).First(&book)
	if rs.Error != nil {
		return book, rs.Error
	}
	return book, nil
}

func (dc *DatabaseController) GetChapterByUnique(bookId, index int) (models.Chapter, error) {
	var chapter models.Chapter

	rs := dc.DB.Preload("Entries."+clause.Associations).Where("book_id = ?", bookId).Where("index = ?", index).First(&chapter)
	if rs.Error != nil {
		return chapter, rs.Error
	}

	return chapter, nil
}

func (dc *DatabaseController) GetChapter(chapterId int) (models.Chapter, error) {
	var chapter models.Chapter

	rs := dc.DB.Preload("Entries."+clause.Associations).Where("id = ?", chapterId).First(&chapter)
	if rs.Error != nil {
		return chapter, rs.Error
	}

	return chapter, nil
}

func (dc *DatabaseController) DeleteChapter(chapterId int) error {
	chapter, err := dc.GetChapter(chapterId)
	if err != nil {
		return err
	}

	entries := chapter.Entries
	err = dc.DB.Model(&chapter).Association("Entries").Clear()
	if err != nil {
		return err
	}

	rs := dc.DB.Unscoped().Delete(&chapter)
	if rs.Error != nil {
		return rs.Error
	}

	for _, entry := range entries {
		rs := dc.DB.Delete(&entry)
		if rs.Error != nil {
			return rs.Error
		}
	}
	return nil
}

func (dc *DatabaseController) GetEntryByUnique(chapterId, index int) (models.Entry, error) {
	var entry models.Entry

	rs := dc.DB.Where("chapter_id = ?", chapterId).Where("index = ?", index).First(&entry)
	if rs.Error != nil {
		return entry, rs.Error
	}

	return entry, nil
}

func (dc *DatabaseController) DropAllTables() error {
	rs := dc.DB.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
	if rs.Error != nil {
		return rs.Error
	}

	log.Println("Drop all tables done!")

	return nil
}
