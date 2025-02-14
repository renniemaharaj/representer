package main

import (
	"github.com/renniemaharaj/representer/internal/ui"
	"github.com/renniemaharaj/representer/internal/utils"
)

func main() {
	// Create a new document
	var document = ui.MyDocument()

	// Transform the document
	transformed := document.Transform()

	print(utils.LintCodeFences(transformed))
	document.ExportMarkup("index.html")
}
