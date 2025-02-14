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
	<Instructions>

    	<h1>Responsive Web Page Generator</h1>  

			<h2>Request</h2>  
				<p>Generate a responsive web page document.</p>  

			<h2>Description</h2>  
				<p>This application generates responsive web pages based on a provided JSON object, which represents a Go struct-based document. Elements will have classes using Tailwind CSS syntax.</p>  

			<h2>Instructions</h2>  
				<ul>  
    				<li>Use the attached JSON object as the data source.</li>  
    				<li>Convert Tailwind CSS utility classes into element's inline styles</li>  
    			<ul>  

	</Instructions>
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
