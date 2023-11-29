package main

import (
	"golang/handlers"
	"net/http"
	"time"
)

func main() {
	// Initialize http service
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		IdleTimeout:  800 * time.Second,
		Handler:      handlers.API(),
	}
	api.ListenAndServe()
}
