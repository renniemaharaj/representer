package elements

// Script represents a script tag in the HTML head.
type Script struct {
	Src     string
	Async   bool
	Defer   bool
	Inline  string
	Content string
}

//This function will return a single script tag.
func MakeScript(src string, async, deferring bool, inline string) *Script {
	return &Script{
		Src:    src,
		Async:  async,
		Defer:  deferring,
		Inline: inline,
	}
}
