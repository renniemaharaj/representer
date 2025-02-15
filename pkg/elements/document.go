package elements

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/renniemaharaj/representer/pkg/utils"
	"google.golang.org/api/option"
)

// HtmlDocument struct represents an entire HTML document.
type Document struct {
	Language string
	Head     Head
	Body     Body
}

// This function creates and returns a blank document skeleton.
func BlankDocument() Document {
	return Document{
		Head: Head{
			Title:   "",
			Metas:   []Meta{},
			Links:   []Link{},
			Scripts: []Script{},
		},
		Body: Body{},
	}
}

// This function generates markup for an HtmlDocument and exports it to the file specified. Export as .html.
func (doc Document) Export(filename string) {
	html := doc.Transform()
	fileName := filename
	err := os.WriteFile(fileName, []byte(utils.LintCodeFences(html)), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}

// MarshalDocument marshals a document to a JSON string
func (doc *Document) Marshal() string {
	documentBytes, err := json.Marshal(doc)

	if err != nil {
		log.Fatalf("Error marshalling document: %v", err)
	}

	return string(documentBytes)
}

// SendMessage sends a message to the generative AI model
func (doc *Document) Transform() string {
	ctx := context.Background()

	apiKey, ok := os.LookupEnv("GEMINI_API_KEY")

	if !ok {
		log.Fatalln("Environment variable GEMINI_API_KEY not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-pro-exp-02-05")

	model.SetTemperature(1)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(utils.GetInstructions()),
		},
	}

	session := model.StartChat()
	session.History = []*genai.Content{}

	log.Print("Transforming document...")

	// Send the document to the model
	resp, err := session.SendMessage(ctx, genai.Text(doc.Marshal()))
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	// Collect response parts
	var messages []string
	for _, part := range resp.Candidates[0].Content.Parts {
		switch v := part.(type) {
		case genai.Text:
			messages = append(messages, string(v))
		default:
			// s.logger.Errorf("unexpected part type: %T", v)
		}
	}
	return utils.LintCodeFences(strings.Join(messages, " "))
}
