package representer

type Style struct {
	Attribute string
	Value     string
}
type StyleBlock struct {
	Selection string
	Styles    []Style
}
