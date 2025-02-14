package representer

import (
	"log"
	"os"
)

// This function generates markup for an HtmlDocument and exports it to the file specified. Export as .html.
func (doc HtmlDocument) ExportMarkup(filename string) {
	html := doc.Transform()
	fileName := filename
	err := os.WriteFile(fileName, []byte(html), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}
