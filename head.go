package representer

// HtmlHead struct represents the entire head section of a document.
type HtmlHead struct {
	Title       string
	Description string
	Keywords    string
	Author      string
	Metas       Metas
	Links       Links
	Scripts     Scripts
}
