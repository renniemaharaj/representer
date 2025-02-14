package ui

import "github.com/renniemaharaj/representer"

var language = "en"
var title = "Document Representer!"
var description = "This is a simple document go representer for building and generating HTML documents."
var author = "Rennie Maharaj"
var keywords = "ai,generated,thewriterco"

func MyDocument() *representer.Document {
	// Create a new document
	var document = representer.BlankDocument()

	// Set our document head
	document.Head = *Head()

	// Set our document language
	document.Language = language

	// Set our document body
	document.Body = *Body()

	// Set our document style
	document.Head.Styles = append(document.Head.Styles, *Style())

	// Return our document
	return &document
}

func Head() *representer.Head {
	head := representer.Head{}

	head.Title = title

	metas := make([]representer.Meta, 0)

	metas = append(metas, *representer.MakeMeta("charset", []string{"UTF-8"}, ""))
	metas = append(metas, *representer.MakeMeta("name", []string{"viewport"}, "width=device-width, initial-scale=1.0"))
	metas = append(metas, *representer.MakeMeta("name", []string{"description"}, description))
	metas = append(metas, *representer.MakeMeta("name", []string{"author"}, author))
	metas = append(metas, *representer.MakeMeta("name", []string{"keywords"}, keywords))
	metas = append(metas, *representer.MakeMeta("charset", []string{"UTF-8"}, ""))

	head.Metas = metas

	return &head
}

func Style() *representer.Style {
	style := representer.Style{}

	style.Selection = "body"
	style.Styles = map[string]string{
		"backgroundColor": "#101211",
		"color":           "white",
	}

	return &style
}
func Body() *representer.Body {
	body := representer.Body{}

	h1 := representer.Element{
		Tag: "h1",
		Attributes: []representer.Attribute{
			{Name: "class", Value: "absolute top-10 left-[50%] translate-x-[-50%] animate-bounce"},
			{Name: "innerHTML", Value: "Go Represent!"},
			{Name: "id", Value: "title"},
		},
		Children: []representer.Element{
			{},
		},
	}

	body.Elements = append(body.Elements, h1)

	return &body
}
