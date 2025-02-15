package utils

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

// RemoveCodeFences removes ```html from the start and ``` from the end of the input string.
func LintCodeFences(input string) string {
	const codeFenceStart = "```html"
	const codeFenceEnd = "```"

	// Trim the starting "```html"
	input = strings.TrimPrefix(input, codeFenceStart)

	// Trim any leading/trailing whitespace or newlines to better detect the ending code fence
	input = strings.TrimSpace(input)

	// Trim the ending "```"
	input = strings.TrimSuffix(input, codeFenceEnd)

	// Trim excess whitespace again
	return strings.TrimSpace(input)
}

// hashFile computes the SHA-256 hash of a file
func HashFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(content)
	return fmt.Sprintf("%x", hash), nil
}
