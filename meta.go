package representer

// MetaTag represents a meta tag in the HTML head.
type Meta struct {
	Attribute string
	Values    []string
	Content   string
}

// Returns a meta tag <meta attribute="value" content="content">. Omit content if necessary
func MakeMeta(attribute string, values []string, content string) *Meta {
	return &Meta{
		Attribute: attribute,
		Values:    values,
		Content:   content,
	}
}
