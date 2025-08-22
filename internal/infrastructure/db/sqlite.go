package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

func Init(dbPath string, models ...interface{}) (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if len(models) > 0 {
		if err := database.AutoMigrate(models...); err != nil {
			return nil, err
		}
	}
	globalDB = database
	log.Printf("sqlite initialized at %s", dbPath)
	return database, nil
}

func DB() *gorm.DB { return globalDB }
