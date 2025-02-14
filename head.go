package representer

// HtmlHead struct represents the entire head section of a document.
type Head struct {
	Title   string
	Metas   []Meta
	Links   []Link
	Styles  []Style
	Scripts []Script
}
