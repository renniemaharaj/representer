package gemini

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"github.com/renniemaharaj/representer/pkg/transformer"
)

func Model(ctx context.Context) (*genai.GenerativeModel, func(), error) {
	apiKey, ok := os.LookupEnv("GEMINI_API_KEY")
	if !ok {
		return nil, nil, fmt.Errorf("environment variable GEMINI_API_KEY not set")
	}

	log.Println("Creating google gemini client...")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, nil, fmt.Errorf("error creating client: %v", err)
	}

	model := client.GenerativeModel("gemini-2.0-pro-exp-02-05")
	model.SetTemperature(1)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(transformer.GetInstructions()),
		},
	}

	return model, func() { client.Close() }, nil
}
