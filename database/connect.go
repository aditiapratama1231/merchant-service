package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres" //for import postgres
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	// host := os.Getenv("DATABASE_HOST")
	username := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	dbURI := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName)

	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		panic("failed to connect to database")
	}
	log.Println("Connected to database")
	return db
}
