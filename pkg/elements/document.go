package elements

import (
	"context"
	"encoding/json"
	"fmt"
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

// This function transforms a Document and exports it to the file specified. Export as .html.
func (doc Document) Export(dist string) {

	response, err := doc.Transform()
	if err != nil {
		log.Fatalf("Error transforming document: %v", err)
	}

	// Create the dist directory if it doesn't exist
	err = os.MkdirAll(dist, 0755)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	err = os.WriteFile(fmt.Sprintf("%v/%v", dist, response.Html.Filename), []byte(response.Html.Content), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	err = os.WriteFile(fmt.Sprintf("%v/%v", dist, response.Css.Filename), []byte(response.Css.Content), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	err = os.WriteFile(fmt.Sprintf("%v/%v", dist, response.Script.Filename), []byte(response.Script.Content), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}

// Marshals a document to a JSON string
func (doc *Document) Marshal() string {
	documentBytes, err := json.Marshal(doc)

	if err != nil {
		log.Fatalf("Error marshalling document: %v", err)
	}

	return string(documentBytes)
}

func UnmarshalResponse(response string) (*utils.ResponseSchema, error) {
	linted := utils.LintCodeFences(&response, "json")

	res := utils.ResponseSchema{}

	err := json.Unmarshal([]byte(*linted), &res)

	if err != nil {
		return &utils.ResponseSchema{}, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &res, nil

}

// Requests a model to transform document
func (doc *Document) Transform() (*utils.ResponseSchema, error) {
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

	//Build the response
	response := strings.Join(messages, " ")

	// Lint the response
	linted := utils.LintCodeFences(&response, "html")

	// Unmarshal the response
	responseStruct, err := UnmarshalResponse(*linted)
	if err != nil {
		return &utils.ResponseSchema{}, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return responseStruct, nil
}
