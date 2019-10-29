package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
	"os"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	log.Print(GetDBDSN())
	db, err := gorm.Open("postgres", GetDBDSN())
	if err != nil {
		panic(err.Error())
	}

	db.DB().SetMaxOpenConns(10)
	return db
}

func GetDBDSN() string {
	return fmt.Sprintf(
		"host=%s sslmode=%s user=%s password=%s dbname=%s",
		os.Getenv("DB_HOST"),
		"disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}

func GetDBConnection() *gorm.DB {
	return db
}
