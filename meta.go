package representer

// MetaTag represents a meta tag in the HTML head.
type MetaTag struct {
	Attribute string
	Values    []string
	Content   string
}

// Returns a meta tag <meta attribute="value" content="content">. Omit content if necessary
func MakeMeta(attribute string, values []string, content string) *MetaTag {
	return &MetaTag{
		Attribute: attribute,
		Values:    values,
		Content:   content,
	}
}

// Type Metas represents the entire meta section of a document head
type Metas struct {
	MetaTags []MetaTag
}

// This function appends a meta tag to the metas struct of a document head.
func (metas *Metas) AppendMeta(metatag *MetaTag) {
	metas.MetaTags = append(metas.MetaTags, *metatag)
}
