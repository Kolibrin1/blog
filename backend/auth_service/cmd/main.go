package main

import (
	"log"
	"net/http"
	"os"

	"auth-service/internal/handlers"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Liveness Probe: сервис работает
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	// Readiness Probe: сервис готов обслуживать запросы
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ready"))
}

func main() {
	r := mux.NewRouter()

	// Добавляем middleware для логирования
	r.Use(loggingMiddleware)

	// Endpoints для Probes
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/ready", readyHandler).Methods("GET")

	// Auth Endpoints
	r.HandleFunc("/login", handlers.Login()).Methods("POST")
	r.HandleFunc("/register", handlers.RegisterUser()).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Auth Service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
