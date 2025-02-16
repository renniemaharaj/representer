package server

import (
	"fmt"
	"log"
	"net/http"
)

// WServer starts an HTTP server that serves the latest document version
func WServer(port string, dist string, chanS chan []byte, chanR chan []byte) {
	mux := http.NewServeMux()

	// Serve all files from dist/
	mux.Handle("/", http.FileServer(http.Dir(dist)))

	// Serve files from static/
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

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
