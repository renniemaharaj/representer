package elements

// Body represents the body of a document.
type Body struct {
	Elements []Element
}

// This function appends an elemement to a document body.
func (body *Body) AppendChild(element *Element) {
	body.Elements = append(body.Elements, *element)
}
