package main

import (
	"fmt"

	"github.com/renniemaharaj/representer/internal/ui"
	"github.com/renniemaharaj/representer/pkg/elements"
	"github.com/renniemaharaj/representer/pkg/server"
)

// The port to run the HTTP server on
var port = "8080"

func JScriptWS() elements.Script {
	return elements.Script{
		Content: fmt.Sprintf(`
			// Create a WebSocket connection
			const ws = new WebSocket("ws://localhost:%v/ws");

			ws.onopen = function(event) {
				console.log("WebSocket connection opened.");

				// Send a message to the server
				ws.send("Hello, server!");

				// Receive messages from the server
				ws.onmessage = function(event) {
					document.querySelector("#title").innerHTML = event.data;
				};
			};
			
		`, port),
	}
}

func main() {
	// Create a new document
	doc := ui.MyDocument()

	// Add the WebSocket script to the document
	doc.Head.Scripts = append(doc.Head.Scripts, JScriptWS())

	// Transform the document to HTML
	doc.Export("index.html")

	// Channels for WebSocket communication
	chanS := make(chan []byte)
	chanR := make(chan []byte)

	// Start the HTTP/WebSocket server
	server.WServer(port, chanS, chanR)

	select {} // Keep the main thread running
}
