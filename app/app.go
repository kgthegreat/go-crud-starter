package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/kgthegreat/meeteffective/config"
	"github.com/kgthegreat/meeteffective/database"
)

type App struct {
	Config   config.Config
	Database *database.SqliteDB
}

func New(cfg config.Config) *App {
	db, err := database.NewSqliteDB(cfg.Sqlite)
	if err != nil {
		log.Fatal(err)
	}

	return &App{cfg, db}
}

func (a *App) Run(r *chi.Mux) {
	//	headersOk := handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "X-Requested-With"})
	//	originsOk := handlers.AllowedOrigins([]string{"*"})
	//	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	port := a.Config.Port
	addr := fmt.Sprintf(":%v", port)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	fmt.Printf("APP is listening on port: %d\n", port)
	log.Fatal(server.ListenAndServe())

}

func (a *App) IsProd() bool {
	return a.Config.Env == "prod"
}
