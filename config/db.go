package config

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var MongoURI string

func init() {

	MongoURI = os.Getenv("MongoURI")

	if MongoURI == "" {
		log.Panic("Error loading MongoURI")
	}

}
