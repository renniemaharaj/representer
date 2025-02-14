package representer

// HtmlDocument struct represents an entire HTML document.
type Document struct {
	Language string
	Head     Head
	Body     Body
}

// This function creates and returns a blank document skeleton.
func BlankDocument() Document {
	return Document{
		Head: Head{
			Title:   "",
			Metas:   []Meta{},
			Links:   []Link{},
			Scripts: []Script{},
		},
		Body: Body{},
	}
}

// This function will run an API call on the HtmlDocument and return its markup representation
func (doc *Document) Transform() string {
	return TransformDocument(*doc)
}
