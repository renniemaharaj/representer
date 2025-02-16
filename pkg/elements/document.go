package elements

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"context"

	"github.com/renniemaharaj/representer/pkg/transformer"
	"github.com/renniemaharaj/representer/pkg/transformer/gemini"

	"github.com/renniemaharaj/representer/pkg/server"
)

// HtmlDocument struct represents an entire HTML document.
type Document struct {
	Language string
	Head     Head
	Body     Body
}

// This function creates and returns a blank document skeleton.
func BlankDocument() *Document {
	return &Document{
		Head: Head{
			Title:   "",
			Metas:   []Meta{},
			Links:   []Link{},
			Scripts: []Script{},
		},
		Body: Body{},
	}
}

// This function builds and serves a document on the specified port and directory.
func (doc *Document) BuildAndServe(port, dist string) (chan []byte, chan []byte) {
	// Transform the document to HTML
	doc.Build(dist)

	// Channels for WebSocket communication
	chanS := make(chan []byte)
	chanR := make(chan []byte)

	// Start the HTTP/WebSocket server
	server.WServer(port, dist, chanS, chanR)

	return chanS, chanR
}

// This function transforms a Document and exports it to the file specified. Export as .html.
func (doc *Document) Build(dist string) error {
	response, err := doc.Transform()
	if err != nil {
		return fmt.Errorf("error transforming document: %v", err)
	}

	if err := os.MkdirAll(dist, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	files := []struct {
		path, content string
	}{
		{fmt.Sprintf("%v/%v", dist, response.Html.Filename), response.Html.Content},
		{fmt.Sprintf("%v/%v", dist, response.Css.Filename), response.Css.Content},
		// {fmt.Sprintf("%v/%v", dist, response.Script.Filename), response.Script.Content},
	}

	for _, file := range files {
		// Skip empty files
		if file.content == "" {
			continue
		}

		// Write the file
		if err := os.WriteFile(file.path, []byte(file.content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", file.path, err)
		}
	}

	return nil
}

// Marshals a document to a JSON string
func (doc *Document) Marshal() ([]byte, error) {
	documentBytes, err := json.Marshal(doc)

	if err != nil {
		return nil, fmt.Errorf("error marshalling document: %v", err)
	}

	return documentBytes, nil

}

// Requests a model to transform document
func (doc *Document) Transform() (*transformer.Schema, error) {
	// Create a new context
	ctx := context.Background()

	// Get the model
	model, cleanup, err := gemini.Model(ctx)
	if err != nil {
		return nil, err
	}
	defer cleanup()

	// Create a new session
	session := gemini.Session{Model: *model}

	// Marshal the document
	docBytes, err := doc.Marshal()
	if err != nil {
		return nil, fmt.Errorf("error marshalling document: %v", err)
	}

	log.Println("Transforming document...")
	return session.Request(string(docBytes), ctx)
}
