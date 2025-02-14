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

// Type Metas represents the entire meta section of a document head
type Metas struct {
	MetaTags []Meta
}

// This function appends a meta tag to the metas struct of a document head.
func (metas *Metas) AppendMeta(metatag *Meta) {
	metas.MetaTags = append(metas.MetaTags, *metatag)
}
