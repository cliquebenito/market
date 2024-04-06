package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"project/config"
	"project/internal/routes"
)

//func init() {
//
//	cfg, err := config.CheckConfig()
//	if err != nil {
//		log.Fatalf("Config error: %s", err)
//	}
//	app.Run(cfg)
//
//}

func main() {
	// Configuration
	cfg, _ := config.CheckConfig()

	s := routes.New(cfg)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}

// https://github.com/ilyakaznacheev/cleanenv/blob/master/example/parse_multiple_files/main.go
