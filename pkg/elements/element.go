package elements

// Element struct represents a document body element
type Element struct {
	Tag        string
	Attributes []Attribute
	Children   []Element
}

// This function appends one element as a child to another
func (parent *Element) AppendChild(element *Element) {
	parent.Children = append(parent.Children, *element)
}

// This function appends attributes to elements.
func (element *Element) AppendAttribute(attribute *Attribute) {
	element.Attributes = append(element.Attributes, *attribute)
}

// This will create and element of type specificed, and take a map for attributes.
func CreateElementByAttributes(tag string, attributes *[]Attribute) Element {
	var element Element = Element{}
	element.Tag = tag
	element.Attributes = *attributes
	return element
}
