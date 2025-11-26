package database

import (
	"MIS/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//locate to database address
	dsn := "host=localhost user=postgres password=930017 dbname=mis port=5432 sslmode=disable"

	//open database

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//auto create
	err = db.AutoMigrate(
		&model.Student{},
		&model.Teacher{},
	)
	if err != nil {
		log.Fatal("Failed to migrate models:", err)
	}

	DB = db

}
