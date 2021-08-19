package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("an error occurred while connecting to the database: ", err)
	}
	err = db.AutoMigrate(&Image{})
	if err != nil {
		log.Fatal("an error occurred while migrating database models: ", err)
	}
	return db
}
