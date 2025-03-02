package main

import (
	"log"
	"net/http"
	"os"

	"posts_service/internal/database"
	"posts_service/internal/handlers"
	"posts_service/internal/middlewares"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Liveness Probe: Service is running
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Readiness Probe: Service is ready to accept traffic
func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ready"))
}

func main() {
	// Подключение к базе данных
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	// Добавляем middleware для логирования
	r.Use(loggingMiddleware)
	r.Use(middlewares.AuthMiddleware)

	// Пробы
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/ready", readyHandler).Methods("GET")

	// Маршруты для постов
	r.HandleFunc("/posts", handlers.CreatePost(db)).Methods("POST")
	r.HandleFunc("/posts", handlers.FetchPosts(db)).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.FetchPostById(db)).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.DeletePost(db)).Methods("DELETE")

	// Маршруты для лайков
	r.HandleFunc("/likes", handlers.ToggleLike(db)).Methods("POST", "DELETE")
	r.HandleFunc("/likes", handlers.GetLikesForPost(db)).Methods("GET")

	// Маршрут для получения постов конкретного пользователя
	r.HandleFunc("/profile/{username}/posts", handlers.FetchUserPosts(db)).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083" // Порт для пост-сервиса
	}

	log.Printf("Posts Service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
