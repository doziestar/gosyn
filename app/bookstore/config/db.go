package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDb() {
	d, err := gorm.Open(postgres.Open("postgres:918273645dozie/bookstore"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
