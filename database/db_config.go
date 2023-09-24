package database

import (
	"fmt"
	"log"

	"github.com/Gilberd-dev/task-5-pbi-btpns-gilberd-nicolas-siboro/helpers"
	"github.com/Gilberd-dev/task-5-pbi-btpns-gilberd-nicolas-siboro/models"
	"gorm.io/driver/mysql" // Import driver MySQL
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() {
	var path string
	stage := helpers.GetAsString("STAGE", "development")

	if stage == "testing" {
		path = "../.env"
	}
	if stage != "testing" {
		path = ".env"
	}

	// comment this line for production ready app (with container)
	helpers.LoadEnv(path)

	dbURI := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		helpers.GetAsString("DB_USER", "root"),
		helpers.GetAsString("DB_HOST", "localhost"),
		helpers.GetAsInt("DB_PORT", 3307),
		helpers.GetAsString("DB_NAME", "mydatabase"),
	)

	db, err = gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func MigrateDB() {
	stage := helpers.GetAsString("STAGE", "development")

	if stage == "development" ||
		stage == "production" {
		db.Debug().AutoMigrate(models.User{}, models.Photo{})
	}
}

func GetDB() *gorm.DB {
	return db
}
