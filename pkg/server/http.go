package server

import (
	"fmt"
	"log"
	"net/http"
)

// WServer starts an HTTP server that serves the latest document version
func WServer(port string, chanS chan []byte, chanR chan []byte) {
	mux := http.NewServeMux()

	// Serve the latest transformed document
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	// WebSocket route
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		WSocket(chanS, chanR)(w, r)
	})

	server := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: mux}

	go func() {
		log.Printf("Server running on port %s...\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
}
