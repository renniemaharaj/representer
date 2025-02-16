package transformer

import (
	"encoding/json"
	"fmt"

	"github.com/renniemaharaj/representer/pkg/utils"
)

type File struct {
	Content  string `json:"content"`
	Filename string `json:"filename"`
}

type Schema struct {
	Html File `json:"html"`
	Css  File `json:"css"`
}

func Unmarshal(resp string) (*Schema, error) {
	linted := utils.LintCodeFences(&resp, "json")
	if linted == nil || *linted == "" {
		return nil, fmt.Errorf("linted response is empty or nil")
	}

	res := Schema{}

	err := json.Unmarshal([]byte(*linted), &res)

	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &res, nil

}
