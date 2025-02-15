package elements

// LinkTag represents a link tag in the document head.
type Link struct {
	Rel  string
	Href string
}

//Returns a new meta tag
func MakeLink(rel, href string) Link {
	return Link{
		Rel:  rel,
		Href: href,
	}
}
