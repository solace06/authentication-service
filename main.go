package main

import (
	"log"

	"github.com/solace06/auth-service/config"
	"github.com/solace06/auth-service/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")
}
