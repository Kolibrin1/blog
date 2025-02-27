package main

import (
	"log"
	"net/http"
	"os"

	"users_service/internal/database"
	"users_service/internal/handlers"

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
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	// Добавляем middleware для логирования
	r.Use(loggingMiddleware)

	// Endpoints for probes
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/ready", readyHandler).Methods("GET")

	// User service endpoints
	r.HandleFunc("/api/users/register", handlers.RegisterUser(db)).Methods("POST")
	r.HandleFunc("/api/users/by_email", handlers.GetUserByEmail(db)).Methods("GET")
	r.HandleFunc("/api/users/by_username", handlers.GetUserByUsername(db)).Methods("GET")
	r.HandleFunc("/api/users/{id:[0-9]+}", handlers.GetUserByID(db)).Methods("GET")
	r.HandleFunc("/api/users/{id:[0-9]+}", handlers.UpdateUser(db)).Methods("PATCH")
	r.HandleFunc("/api/users/{id:[0-9]+}/password", handlers.UpdateUserPassword(db)).Methods("PATCH")
	r.HandleFunc("/api/users/{id:[0-9]+}", handlers.DeleteUser(db)).Methods("DELETE")
	r.HandleFunc("/api/users", handlers.ListUsers(db)).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}
	log.Printf("Users Service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
