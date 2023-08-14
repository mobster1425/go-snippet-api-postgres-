package initialize

import (
	"log"

	"feyin/go-restapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"github.com/joho/godotenv"
	"github.com/thedevsaddam/renderer"
)
/*
Initialize Function (Init() (*gorm.DB,*renderer.Render)):

Initializes the database connection and renderer.
Reads database connection URI from environment variables.
Performs database auto migration to create tables if they don't exist.
Returns the database connection and renderer instances.
*/
func Init() (*gorm.DB,*renderer.Render) {
	rnd := renderer.New()
	
	godotenv.Load();

	uri := os.Getenv("dbURL")
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Snippet{})

	return db,rnd
}