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
	sb := strings.Builder{}
	for _, part := range resp.Candidates[0].Content.Parts {
		sb.WriteString(string(part.(genai.Text)))
	}

	//Build the response
	response := sb.String()

	// Lint the response
	linted := utils.LintCodeFences(&response, "html")

	// Unmarshal the response
	responseStruct, err := transformer.Unmarshal(*linted)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return responseStruct, nil
}
