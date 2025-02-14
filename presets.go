package representer

// CreateElement creates an element with specified tag, class, id, onclick, and content
func CreateElement(tag, class, id, onclick, content string) Element {
	attrs := []Attribute{}

	if class != "" {
		attrs = append(attrs, Attribute{Name: "class", Value: class})
	}
	if id != "" {
		attrs = append(attrs, Attribute{Name: "id", Value: id})
	}
	if onclick != "" {
		attrs = append(attrs, Attribute{Name: "onclick", Value: onclick})
	}
	if content != "" {
		attrs = append(attrs, Attribute{Name: "content", Value: onclick})
	}

	return CreateElementByAttributes(tag, &attrs)
}
