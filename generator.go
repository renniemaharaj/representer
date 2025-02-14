package representer

import (
	"context"
	"encoding/json"
	"log"
	"os"

	// "os"
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
				<p>This application generates responsive web pages based on a provided JSON object, which represents a Go struct-based document. Elements utilize tailwind utility classes in their class attribute.</p>  

			<h2>Instructions</h2>  
				<ul>  
    				<li>Use the attached JSON object as the data source.</li>  
    				<li>Convert Tailwind utility classes to it's valid equivalent vanilla CSS selector and style body.</li>
					<li>If no style tags are included, please insert one and generate vanilla CSS there</li>
					<li>This is an application component, do not generate text or explanations<li>
					<li>Do not double escape CSS selectors like : <code>.left-\\[50\\%\\]</code> This is invalid CSS<li>
					<li>You must escape once like: <code>.left-\[50\%\]</code><li>
					<li>Please make sure that transform-based animations do not break other transformations. You may need to couple transformations to account for transformation styles outside keyframes<li>  
    			<ul>  

	</Instructions>
	`
}

// MarshalDocument marshals a document to a JSON string
func MarshalDocument(document Document) string {
	documentBytes, err := json.Marshal(document)

	if err != nil {
		log.Fatalf("Error marshalling document: %v", err)
	}

	return string(documentBytes)
}

// SendMessage sends a message to the generative AI model
func TransformDocument(document Document) string {
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

	log.Print("Transforming document...")

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
	return LintCodeFences(strings.Join(messages, " "))
}
