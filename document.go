package representer

// HtmlDocument struct represents an entire HTML document.
type HtmlDocument struct {
	Language string
	Head     HtmlHead
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
func BlankDocument() HtmlDocument {
	return HtmlDocument{
		Head: HtmlHead{
			Title:   "",
			Metas:   Metas{},
			Links:   Links{},
			Scripts: Scripts{},
		},
		Body: Body{},
	}
}

// This function will run an API call on the HtmlDocument and return its markup representation
func (doc *HtmlDocument) Transform() string {
	return TransformDocument(*doc)
}
