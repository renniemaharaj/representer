package ui

import "github.com/renniemaharaj/representer"

var language = "en"
var title = "Document Representer!"
var description = "This is a simple document representer using GoLang for generating HTML documents."
var author = "Rennie Maharaj"
var keywords = "ai,generated,thewriterco"

func MyDocument() *representer.Document {
	// Create a new document
	var document = representer.BlankDocument()

	// Set the document head
	document.Head = *Head()

	// Set the document language and title
	document.Language = language
	document.Head.Title = title

	// Append the body to the document
	document.Body = *Body()

	return &document
}

func Head() *representer.Head {
	head := representer.Head{}

	var metas = &representer.Metas{}

	metas.AppendMeta(representer.MakeMeta("charset", []string{"UTF-8"}, ""))
	metas.AppendMeta(representer.MakeMeta("name", []string{"viewport"}, "width=device-width, initial-scale=1.0"))
	metas.AppendMeta(representer.MakeMeta("name", []string{"description"}, description))
	metas.AppendMeta(representer.MakeMeta("name", []string{"author"}, author))
	metas.AppendMeta(representer.MakeMeta("name", []string{"keywords"}, keywords))

	head.Metas = *metas

	return &head
}

func Body() *representer.Body {
	body := representer.Body{}

	header := representer.Element{
		Tag: "h1",
		Attributes: []representer.Attribute{
			{Name: "class", Value: "absolute top-10 left-[50%] translate-x-[-50%] animate-bounce"},
			{Name: "innerHTML", Value: "Representer"},
		},
		Children: []representer.Element{
			{},
		},
	}

	body.Elements = append(body.Elements, header)

	return &body
}
