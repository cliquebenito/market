package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"project/config"
	"project/internal/app"
)

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
