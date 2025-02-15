package server

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/renniemaharaj/representer/pkg/utils"
)

// WatchFile detects changes in index.html and sends updates via WebSocket
func WatchFile(filePath string, chanS chan []byte, mu *sync.Mutex) {
	lastHash := ""

	for {
		time.Sleep(1 * time.Second) // Polling interval

		mu.Lock()
		hash, err := utils.HashFile(filePath)
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
