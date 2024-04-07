package database

import (
	"log"
	"testing"

	"github.com/go_learning/gorm_demo/database/models"
	"github.com/stretchr/testify/require"
)

func TestCreateUpdate(t *testing.T) {
	db := New()
	db.DropAllTables()

	db = New()

	book := models.Book{
		Name: "Book1",
		Chapters: []models.Chapter{
			{
				Name:  "Chapter 1",
				Index: 1,
				Entries: []models.Entry{
					{
						Name:  "Entry 1.1",
						Index: 1,
					},
					{
						Name:  "Entry 1.2",
						Index: 2,
					},
					{
						Name:  "Entry 1.3",
						Index: 3,
					},
				},
			},
			{
				Name:  "Chapter 2",
				Index: 2,
				Entries: []models.Entry{
					{
						Name:  "Entry 2.1",
						Index: 1,
					},
					{
						Name:  "Entry 2.2",
						Index: 2,
					},
					{
						Name:  "Entry 2.3",
						Index: 3,
					},
				},
			},
			{
				Name:  "Chapter 3",
				Index: 3,
				Entries: []models.Entry{
					{
						Name:  "Entry 3.1",
						Index: 1,
					},
					{
						Name:  "Entry 3.2",
						Index: 2,
					},
					{
						Name:  "Entry 3.3",
						Index: 3,
					},
				},
			},
		},
	}

	err := db.Insert(book, NONE)
	require.NoError(t, err)

	book, err = db.GetBookByUnique(book.Name)
	require.NoError(t, err)

	log.Printf("book %v", book)

	var deleteChapter models.Chapter
	for i := 0; i < len(book.Chapters); i++ {
		if book.Chapters[i].Index == 1 {
			deleteChapter = book.Chapters[i]
			book.Chapters = append(book.Chapters[:i], book.Chapters[i+1:]...)
			log.Printf("book after remove 1 item : %v", book)
			i--
		} else if book.Chapters[i].Index > 1 {
			book.Chapters[i].Index = book.Chapters[i].Index - 1
			log.Printf("chpater i:%d, index after : %d", i, book.Chapters[i].Index)
		}
	}

	log.Printf("book after delete %v", book)

	err = db.DeleteChapter(deleteChapter.ID)
	require.NoError(t, err)

	err = db.Insert(book, ASSOCIATIONS)
	require.NoError(t, err)

	book, err = db.GetBookByUnique(book.Name)
	require.NoError(t, err)
	require.Equal(t, 2, len(book.Chapters))

	require.NotEqual(t, book.Chapters[0].Index, book.Chapters[1].Index)
	require.Equal(t, 3, book.Chapters[0].Index+book.Chapters[1].Index)
}
