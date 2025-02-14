package representer

import (
	"log"
	"os"
	"strings"
)

// This function generates markup for an HtmlDocument and exports it to the file specified. Export as .html.
func (doc HtmlDocument) ExportMarkup(filename string) {
	html := doc.Transform()
	fileName := filename
	err := os.WriteFile(fileName, []byte(LintCodeFences(html)), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}

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
