package utils

import "strings"

// RemoveCodeFences removes ```html from the start and ``` from the end of the input string.
func LintCodeFences(input string) string {
	const codeFenceStart = "```html"
	const codeFenceEnd = "```"

	// Trim the starting "```html"
	input = strings.TrimPrefix(input, codeFenceStart)

	// Trim any leading/trailing whitespace or newlines to better detect the ending code fence
	input = strings.TrimSpace(input)

	// Trim the ending "```"
	input = strings.TrimSuffix(input, codeFenceEnd)

	// Trim excess whitespace again
	return strings.TrimSpace(input)
}
