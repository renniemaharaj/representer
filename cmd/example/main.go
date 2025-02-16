package main

import (
	"github.com/renniemaharaj/representer/internal/ui"
	"github.com/renniemaharaj/representer/pkg/elements"
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

	// Build and serve the document

	// chanS, chanR := doc.BuildAndServe(port, dist)
	doc.BuildAndServe(port, dist)

	select {} // Keep the main thread running
}
