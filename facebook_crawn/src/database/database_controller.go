package database

import (
	"log"

	"github.com/kimvnhung/go_learning/facebook_crawn/src/models"
	"github.com/kimvnhung/go_learning/facebook_crawn/src/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseController struct {
	DB           *gorm.DB
	DataFileName string
}

func New() *DatabaseController {

	p := &DatabaseController{
		DB:           prepareDatabase("facebook_crawn"),
		DataFileName: "facebook_crawn",
	}

	return p
}

func prepareDatabase(databaseName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sample.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.House{}, &models.News{})

	return db
}

func (p *DatabaseController) Insert(news *models.News) error {
	rs := p.DB.Create(news)
	if rs.Error != nil {
		return rs.Error
	}

	house, err := utils.ExtractHouse(news.Content)
	if err != nil {
		return err
	}

	rs = p.DB.Create(house)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

func (p *DatabaseController) GetNews() ([]models.News, error) {
	var news []models.News
	rs := p.DB.Preload("HouseDetail").Find(&news)
	if rs.Error != nil {
		return nil, rs.Error
	}
	return news, nil
}

func (p *DatabaseController) GetHouses(filterField []models.HouseFilterField, filterType []models.HouseFilterType, filterValue []string) ([]models.House, error) {
	var houses []models.House
	var rs *gorm.DB
	if len(filterField) == 0 {
		rs = p.DB.Find(&houses)
	} else {
		// Loop through filterField, filterType, filterValue
		// Add filter condition to rs
		for i := 0; i < len(filterField); i++ {
			// switch case for filterType[i]
			// Add filter condition to rs
			switch filterType[i] {
			case models.Equal:
				rs = p.DB.Where(string(filterField[i])+" = ?", filterValue[i]).Find(&houses)
			case models.NotEqual:
				rs = p.DB.Where(string(filterField[i])+" != ?", filterValue[i]).Find(&houses)
			case models.Contain:
				rs = p.DB.Where(string(filterField[i])+" LIKE ?", "%"+filterValue[i]+"%").Find(&houses)
			case models.NotContain:
				rs = p.DB.Where(string(filterField[i])+" NOT LIKE ?", "%"+filterValue[i]+"%").Find(&houses)
			case models.Greater:
				rs = p.DB.Where(string(filterField[i])+" > ?", filterValue[i]).Find(&houses)
			case models.GreaterOrEqual:
				rs = p.DB.Where(string(filterField[i])+" >= ?", filterValue[i]).Find(&houses)
			case models.Less:
				rs = p.DB.Where(string(filterField[i])+" < ?", filterValue[i]).Find(&houses)
			case models.LessOrEqual:
				rs = p.DB.Where(string(filterField[i])+" <= ?", filterValue[i]).Find(&houses)
			default:
				// do nothing
			}
		}
	}

	if rs.Error != nil {
		return nil, rs.Error
	}
	return houses, nil
}
