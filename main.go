package main

import (
	"log"

	"github.com/kgthegreat/meeteffective/app"
	"github.com/kgthegreat/meeteffective/config"
	"github.com/kgthegreat/meeteffective/routes"
)

func main() {
	cfg, err := config.New("config/app.json")
	if err != nil {
		log.Fatal(err)
	}
	app := app.New(cfg)
	router := routes.NewRouter(app)
	app.Run(router)
}
