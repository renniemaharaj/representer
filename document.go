package representer

// HtmlDocument struct represents an entire HTML document.
type Document struct {
	Language string
	Head     Head
	Body     Body
}

// Body represents the body of a document.
type Body struct {
	Elements []Element
}

// This function appends an elemement to a document body.
func (body *Body) AppendChild(element *Element) {
	body.Elements = append(body.Elements, *element)
}

// This function creates and returns a blank document skeleton.
func BlankDocument() Document {
	return Document{
		Head: Head{
			Title:   "",
			Metas:   Metas{},
			Links:   Links{},
			Scripts: Scripts{},
		},
		Body: Body{},
	}
}

// This function will run an API call on the HtmlDocument and return its markup representation
func (doc *Document) Transform() string {
	return TransformDocument(*doc)
}
