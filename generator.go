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
    			<ul>  
        			<li>Convert <code>left-[50%]</code> → <code>left: 50%;</code></li> 
					<li>Translate selectors with escaped values: <code>left-[50%]</code> → <code>.left-\[50\%\]</code> (NOTICE ONE LEVEL OF ESCAPE | DOUBLE ESCAPING '//' WILL BREAK CSS SYNTAX) </li>
					<li>DO NOT DOUBLE ESCAPE IN SELECTORS!! DO NOT DOUBLE ESCAPE!! .translate-x-\\[-50\\%\\] <- IS NOT VALID CSS, SHOULD BE -> .translate-x-\[-50\%\]<li>
					<li>Translate all selectors appropriately as the above example</li>
        			<li>Convert <code>h-[calc(100vh-4rem)]</code> → <code>height: calc(100vh - 4rem);</code></li>  
    			</ul>  
    				<li>Handle negative values appropriately:</li>  
    			<ul>  
        			<li>Convert <code>-translate-x-1/2</code> → <code>transform: translateX(-50%);</code></li>  
        			<li>Convert <code>-top-4</code> → <code>top: -1rem;</code> (assuming <code>1rem = 4px</code>).</li>  
    			</ul>  
    				<li>Convert shorthand properties into full CSS declarations:</li>  
    			<ul>  
        			<li><code>m-4</code> → <code>margin: 1rem;</code></li>  
        			<li><code>grid-cols-3</code> → <code>grid-template-columns: repeat(3, minmax(0, 1fr));</code></li>  
    			</ul>  
    				<li>Translate conditional classes (<code>hover:</code>, <code>md:</code>, <code>lg:</code>) into valid CSS:</li>  
    			<ul>  
        			<li><code>hover:bg-red-500</code> →  
            		<pre><code>.hover\:bg-red-500:hover { background-color: #ef4444; }</code></pre>  
        			</li>  
        			<li><code>md:flex</code> →  
            		<pre><code>@media (min-width: 768px) { .md\:flex { display: flex; } }</code></pre>  
        			</li>  
    			</ul>  
    				<li>Ensure Tailwind's default units (<code>rem</code>, <code>px</code>, etc.) are respected:</li>  
    			<ul>  
        			<li><code>text-4xl</code> → <code>font-size: 2.25rem;</code></li>  
        			<li><code>gap-2</code> → <code>gap: 0.5rem;</code></li>  
    			</ul>  
    				<li>Convert Tailwind color names to their hex values:</li>  
    			<ul>  
        			<li><code>bg-blue-500</code> → <code>background-color: #3b82f6;</code></li>  
        			<li><code>text-gray-700</code> → <code>color: #374151;</code></li>  
    			</ul>  
    				<li>Handle font weights, styles, and spacing properties:</li>  
    			<ul>  
        			<li><code>font-bold</code> → <code>font-weight: 700;</code></li>  
        			<li><code>italic</code> → <code>font-style: italic;</code></li>  
        			<li><code>leading-tight</code> → <code>line-height: 1.25;</code></li>  
        			<li><code>tracking-wide</code> → <code>letter-spacing: 0.025em;</code></li>  
    			</ul>  
    				<li>Ensure CSS rules do not conflict when multiple utility classes are applied to the same element.</li>  
    				<li>Place all styles within a <code>&lt;style&gt;</code> tag in the <code>&lt;head&gt;</code> section, avoiding inline styles.</li>  
    				<li>If no <code>&lt;style&gt;</code> tag exists, create one.</li>  
    				<li>Respond only with the complete HTML document—no additional text or explanations.</li>  
				</ul>  

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
