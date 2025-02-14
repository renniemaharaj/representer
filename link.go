package representer

//Links struct represents the entire links section of a document head.
type Links struct {
	LinkTags []Link
}

// LinkTag represents a link tag in the document head.
type Link struct {
	Rel  string
	Href string
}

// This functions appends a link to a document head links struct
func (links *Links) AppendLink(link *Link) {
	links.LinkTags = append(links.LinkTags, *link)
}

//Returns a new meta tag
func MakeLink(rel, href string) Link {
	return Link{
		Rel:  rel,
		Href: href,
	}
}
