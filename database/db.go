package database

import (
	"fmt"
	"log"
	"tugas10/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var (
	host    = "localhost"
	user	= "postgres"
	password = "1234"
	port    = "5432"
	dbname =  "tugas10"
	db *gorm.DB
	err error
)

func StartDB(){
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database",err)
	}

	fmt.Println("database connected")
	db.AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}