package representer

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// GetInstructions returns the system instructions
func GetInstructions() string {
	return `
	Request: Generate a responsive web page document.

	Description:
	This application generates responsive web pages based on a provided JSON object, 
	which represents a Go struct-based document.

	Instructions:
	- Use the attached JSON object as the data source.
	- Generate only the HTML content of the web page.
	- Use inline CSS for styling.
	- Do not include any additional text or explanations in the response.
	`
}

// MarshalDocument marshals a document to a JSON string
func MarshalDocument(document HtmlDocument) string {
	documentBytes, err := json.Marshal(document)

	if err != nil {
		log.Fatalf("Error marshalling document: %v", err)
	}

	return string(documentBytes)
}

// SendMessage sends a message to the generative AI model
func TransformDocument(document HtmlDocument) string {
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
			genai.Text(GetInstructions()),
		},
	}

	session := model.StartChat()
	session.History = []*genai.Content{}

	// Send the document to the model
	resp, err := session.SendMessage(ctx, genai.Text(MarshalDocument(document)))
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
	return strings.Join(messages, " ")
}
