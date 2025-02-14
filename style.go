package representer

type Style struct {
	Selection string
	Styles    map[string]string
}

func MakeStyle(selection string, styles map[string]string) *Style {
	return &Style{
		Selection: selection,
		Styles:    styles,
	}
}
