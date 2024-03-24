package db

import (
	"api1/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "postgres"
)

var DB *gorm.DB

func ConnectToPostgresDB() *gorm.DB {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	postgresURL := postgres.Open(url)

	db, err := gorm.Open(postgresURL, &gorm.Config{})

	if err != nil {
		log.Fatalln(err)

	}
	db.AutoMigrate(&model.User1{})
	DB = db
	return db
}
