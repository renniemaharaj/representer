package main

import "representer/representer"

func myDocument() *representer.HtmlDocument {
	// Create a new document
	var document = representer.BlankDocument()

	// Add a meta tag to the document head
	var metas = &representer.Metas{}

	// Append a charset meta tag
	metas.AppendMeta(representer.MakeMeta("charset", []string{"UTF-8"}, ""))

	// Append a viewport meta tag
	metas.AppendMeta(representer.MakeMeta("name", []string{"viewport"}, "width=device-width, initial-scale=1.0"))

	// Set the document head metas
	document.Head.Metas = *metas

	// Set the document language
	document.Language = "en"

	// Set the document title
	document.Head.Title = "Document Representor!"

	// Set the document description, author and keywords
	document.Head.Description = "This has been transformed by the document representor!"
	document.Head.Author = "Rennie Maharaj"
	document.Head.Keywords = "ai,generated,thewriterco"

	return &document
}

func main() {
	var document = myDocument()

	transformed := document.Transform()

	print(transformed)
}
