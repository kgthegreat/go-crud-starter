package main

import (
	"log"

	"github.com/kgthegreat/go-crud-starter/app"
	"github.com/kgthegreat/go-crud-starter/config"
	"github.com/kgthegreat/go-crud-starter/routes"
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
