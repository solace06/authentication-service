package main

import (
	"log"
	"net/http"

	"github.com/solace06/auth-service/api"
	"github.com/solace06/auth-service/config"
	"github.com/solace06/auth-service/database"
	"github.com/solace06/auth-service/internal/user"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = user.NewScope(db)
	if err != nil {
		log.Fatal(err)
	}

	router := api.NewRouter()

	http.ListenAndServe(cfg.ServerPort, router)
}
