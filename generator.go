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
    				<li>Convert Tailwind CSS utility classes into equivalent vanilla CSS.</li>  
    				<li>Properly interpret Tailwind's bracket notation, escaped values, and arbitrary values:</li> 
					<li>Hello, please do not double escape when you're generating vanilla CSS selectors as it will break CSS syntax.
					<li>For example, the following Tailwind class selectors:</li> 
					<pre>This is invalid CSS selector due to double escape characters: <code>.translate-x-\\[-50\\%\\]</code></pre>
					<pre>This is valid CSS selector due to single escape characters: <code>.translate-x-\[-50\%\]</code></pre>
					<li><code>\\</code> is an escape character in Tailwind CSS bracket notation. It is used to escape special characters like <code>[</code> and <code>]</code>.</li>
					<li>Please ensure your output is runnable as this document will be served immediately</li>
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
