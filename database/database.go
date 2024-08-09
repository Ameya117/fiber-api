package database

import (
	"log"
	"os"

	"github.com/Ameya117/fiber-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the database \n", err.Error())
		os.Exit(2)
		// panic("Could not connect to the database")
	}

	log.Println("Connected to the database SUCCESSFULLY")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations") // migration means creating tables in the database
	db.AutoMigrate(&models.User{}, models.Product{}, models.Order{})

	Database = DbInstance{Db: db}

}
