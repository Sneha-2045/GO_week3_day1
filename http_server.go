package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware to log request details
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf(
			"Method=%s Path=%s RemoteAddr=%s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
		)

		next.ServeHTTP(w, r)

		log.Printf(
			"Completed in %v",
			time.Since(start),
		)
	})
}

// Home Handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Internship HTTP Server")
}

// Health Check Endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Server is Healthy")
}

// About Endpoint
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This server is built using Go net/http package")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/about", aboutHandler)

	loggedMux := loggingMiddleware(mux)

	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}