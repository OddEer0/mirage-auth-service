package main

import (
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/config"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	router := chi.NewRouter()
	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte("Hello"))
	})
	err := http.ListenAndServe(cfg.Server.Address, router)
	if err != nil {
		return
	}
}
