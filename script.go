package representer

//Scripts tag represents the entire script section of a document head
type Scripts struct {
	ScriptTags []Script
}

// Script represents a script tag in the HTML head.
type Script struct {
	Src    string
	Async  bool
	Defer  bool
	Inline string
}

// This function will append a script struct to a document head script struct.
func (scripts *Scripts) AppendScript(script *Script) {
	scripts.ScriptTags = append(scripts.ScriptTags, *script)
}

//This function will return a single script tag.
func MakeScript(src string, async, deferring bool, inline string) Script {
	return Script{
		Src:    src,
		Async:  async,
		Defer:  deferring,
		Inline: inline,
	}
}
