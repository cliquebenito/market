package routes

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"project/config"
)

type ApiConfig struct {
	cfg    *config.Config
	router chi.Router
}

func New(config *config.Config) *ApiConfig {
	s := &ApiConfig{
		cfg:    config,
		router: chi.NewRouter(),
	}

	return s
}

func (s *ApiConfig) Start() error {
	fmt.Println("Starting server")

	s.configRoute()

	return http.ListenAndServe(s.cfg.Api.Port, s.router)
}

func (s *ApiConfig) configRoute() {

	s.router.Get("/products", AddProduct)
	s.router.Get("/get", GetProducts)

}
