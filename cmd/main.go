package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
	"project/config"
	"project/internal/app"
	"project/internal/routes"
)

func init() {

	r := chi.NewRouter()
	r.Get("/products", routes.AddProduct)
	log.Println("server started")
	http.ListenAndServe(":8080", r)

}

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)

}

// https://github.com/ilyakaznacheev/cleanenv/blob/master/example/parse_multiple_files/main.go
