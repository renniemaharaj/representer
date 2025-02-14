package representer

//Links struct represents the entire links section of a document head.
type Links struct {
	LinkTags []LinkTag
}

// LinkTag represents a link tag in the document head.
type LinkTag struct {
	Rel  string
	Href string
}

// This functions appends a link to a document head links struct
func (links *Links) AppendLink(link *LinkTag) {
	links.LinkTags = append(links.LinkTags, *link)
}

//Returns a new meta tag
func MakeLink(rel, href string) LinkTag {
	return LinkTag{
		Rel:  rel,
		Href: href,
	}
}
