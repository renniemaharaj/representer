package gemini

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/renniemaharaj/representer/pkg/transformer"
	"github.com/renniemaharaj/representer/pkg/utils"
)

type Session struct {
	Model genai.GenerativeModel
}

func (s *Session) Request(doc string, ctx context.Context) (*transformer.Schema, error) {
	session := s.Model.StartChat()
	session.History = []*genai.Content{}

	// Send the document to the model
	resp, err := session.SendMessage(ctx, genai.Text(doc))
	if err != nil {
		return nil, fmt.Errorf("error sending message: %v", err)
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
	responseStruct, err := transformer.Unmarshal(*linted)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return responseStruct, nil
}
