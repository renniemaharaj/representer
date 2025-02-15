package utils

// "os"

// GetInstructions returns the system instructions
func GetInstructions() string {
	return `
	<Instructions>

    	<h1>Responsive Web Page Generator</h1>  

			<h2>Request</h2>  
				<p>Generate a responsive web page document.</p>  

			<h2>Description</h2>  
				<p>This application generates responsive web pages based on a provided JSON object, which represents a Go struct-based document. Elements utilize tailwind utility classes in their class attribute.</p>  

			<h2>Instructions</h2>  
				<ul>  
    				<li>Use the attached JSON object as the data source.</li>  
    				<li>Convert Tailwind utility classes to it's valid equivalent vanilla CSS selector and style body.</li>
					<li>If no style tags are included, please insert one and generate vanilla CSS there</li>
					<li>This is an application component, do not generate text or explanations<li>
					<li>Do not double escape CSS selectors like : <code>.left-\\[50\\%\\]</code> This is invalid CSS<li>
					<li>You must escape once like: <code>.left-\[50\%\]</code><li>
					<li>Please make sure that transform-based animations do not break other transformations. You may need to couple transformations to account for transformation styles outside keyframes<li>  
    			<ul>  

	</Instructions>
	`
}
