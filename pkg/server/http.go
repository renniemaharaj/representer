package server

import (
	"fmt"
	"log"
	"net/http"
)

// WServer starts an HTTP server that serves the latest document version
func WServer(port string, dist string, chanS chan []byte, chanR chan []byte) {
	mux := http.NewServeMux()

	// Serve the latest transformed document
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fmt.Sprintf("%s/index.html", dist))
	})

	// Serve static files like styles.css and script.js
	mux.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fmt.Sprintf("%s/styles.css", dist))
	})

	mux.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fmt.Sprintf("%s/script.js", dist))
	})

	// WebSocket route
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		WSocket(chanS, chanR)(w, r)
	})

	server := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: mux}

	go func() {
		log.Printf("Server running on port %s...\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("\nServer error: %v", err)
		}
	}()
}
