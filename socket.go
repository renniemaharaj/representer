package representer

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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

// WSocket handles WebSocket connections
func WSocket(chanS chan []byte, chanR chan []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}
		defer conn.Close()

		done := make(chan struct{})

		// Read loop (for receiving messages from client)
		go func() {
			defer close(done)
			for {
				_, msg, err := conn.ReadMessage()
				if err != nil {
					log.Println("Read error:", err)
					return
				}
				chanR <- msg // Pass message to main app logic
			}
		}()

		// Write loop (for sending updates)
		go func() {
			for {
				select {
				case response := <-chanS:
					if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
						log.Println("Write error:", err)
						return
					}
				case <-done:
					return
				}
			}
		}()
	}
}

// WatchFile detects changes in index.html and sends updates via WebSocket
func WatchFile(filePath string, chanS chan []byte, mu *sync.Mutex) {
	lastHash := ""

	for {
		time.Sleep(1 * time.Second) // Polling interval

		mu.Lock()
		hash, err := hashFile(filePath)
		mu.Unlock()
		if err != nil {
			log.Println("Error hashing file:", err)
			continue
		}

		if hash != lastHash {
			log.Println("index.html changed, sending update...")
			content, err := os.ReadFile(filePath)
			if err != nil {
				log.Println("Error reading file:", err)
				continue
			}
			chanS <- content
			lastHash = hash
		}
	}
}

// hashFile computes the SHA-256 hash of a file
func hashFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(content)
	return fmt.Sprintf("%x", hash), nil
}
