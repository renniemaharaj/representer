package main

import (
	// "fmt"

	"github.com/renniemaharaj/representer/internal/ui"
	// "github.com/renniemaharaj/representer/pkg/elements"
	"github.com/renniemaharaj/representer/pkg/elements"
	"github.com/renniemaharaj/representer/pkg/server"
)

// The port to run the HTTP server on
var port = "8080"

// The directory to export the document to
var dist = "dist"

func main() {
	// Create a new document
	doc := ui.MyDocument()

	doc.Head.Links = append(doc.Head.Links, elements.Link{
		Rel:  "stylesheet",
		Href: "/static/styles.css",
	})

	// Add the WebSocket script to the document
	doc.Head.Scripts = append(doc.Head.Scripts, elements.Script{
		Src: "/static/script.js",
	})

	// Transform the document to HTML
	doc.Build(dist)

	// Channels for WebSocket communication
	chanS := make(chan []byte)
	chanR := make(chan []byte)

	// Start the HTTP/WebSocket server
	server.WServer(port, dist, chanS, chanR)

	select {} // Keep the main thread running
}
