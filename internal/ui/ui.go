package ui

import "github.com/renniemaharaj/representer/pkg/elements"

var language = "en"
var title = "Document Representer!"
var description = "This is a simple document go representer for building and generating HTML documents."
var author = "Rennie Maharaj"
var keywords = "ai,generated,thewriterco"

func MyDocument() *elements.Document {
	// Create a new document
	var document = elements.BlankDocument()

	// Set our document head
	document.Head = *Head()

	// Set our document language
	document.Language = language

	// Set our document body
	document.Body = *Body()

	// Set our document style
	document.Head.Styles = append(document.Head.Styles, *Style())

	// Return our document
	return document
}

func Head() *elements.Head {
	head := elements.Head{}

	head.Title = title

	metas := make([]elements.Meta, 0)

	metas = append(metas, *elements.MakeMeta("charset", []string{"UTF-8"}, ""))
	metas = append(metas, *elements.MakeMeta("name", []string{"viewport"}, "width=device-width, initial-scale=1.0"))
	metas = append(metas, *elements.MakeMeta("name", []string{"description"}, description))
	metas = append(metas, *elements.MakeMeta("name", []string{"author"}, author))
	metas = append(metas, *elements.MakeMeta("name", []string{"keywords"}, keywords))
	metas = append(metas, *elements.MakeMeta("charset", []string{"UTF-8"}, ""))

	head.Metas = metas

	return &head
}

func Style() *elements.Style {
	style := elements.Style{}

	style.Selection = "body"
	style.Styles = map[string]string{
		"color": "white",
	}

	return &style
}

func Body() *elements.Body {
	body := elements.Body{}

	h1 := elements.Element{
		Tag: "h1",
		Attributes: []elements.Attribute{
			{Name: "class", Value: "absolute top-10 left-[50%] translate-x-[-50%] animate-bounce"},
			{Name: "innerHTML", Value: "Go Represent!"},
			{Name: "id", Value: "title"},
		},
		Children: []elements.Element{
			{},
		},
	}

	body.Elements = append(body.Elements, h1)

	return &body
}
