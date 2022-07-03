package gormAdapter

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(dsn string) (db *gorm.DB) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to open db", err)
	}

	return
}

func RunMigrations(db *gorm.DB, dbSchemas ...interface{}) {
	db.AutoMigrate(dbSchemas...)
}
