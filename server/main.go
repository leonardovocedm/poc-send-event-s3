package main

import (
	"net/http"
	"os"

	"github.com/Valgard/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test - POST\n"))
		w.WriteHeader(http.StatusCreated)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test - Get\n"))
		w.WriteHeader(http.StatusCreated)
	})

	certFile := os.Getenv("PATH_SERVER_CERT")
	keyFile := os.Getenv("PATH_SERVER_KEY")

	err = http.ListenAndServeTLS(":8080", certFile, keyFile, r)
	if err != nil {
		panic(err)
	}
}
