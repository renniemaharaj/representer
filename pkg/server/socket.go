package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

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
