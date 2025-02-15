package transformer

// GetInstructions returns the system instructions
func GetInstructions() string {
	return `
	<Instructions>

    	<h1>Responsive Web Page Generator</h1>  

		<h2>Request</h2>  
			<p>Generate a responsive web page document with separate HTML, CSS, and JavaScript files.</p>  

		<h2>Description</h2>  
			<p>This application generates responsive web pages based on a provided JSON object, which represents a Go struct-based document. Elements utilize Tailwind utility classes in their class attribute.</p>  

		<h2>Response Format</h2>
			<p>The response must be a JSON object containing three files: HTML, CSS, and JavaScript. Each file should be structured as follows:</p>
			<pre>{"html": {"content": "HTML_CONTENT", "filename": "index.html"}, "css": {"content": "CSS_CONTENT", "filename": "styles.css"}, "script": {"content": "JS_CONTENT", "filename": "script.js"}}</pre>

		<h2>Instructions</h2>  
			<ul>  
    			<li>Use the attached JSON object as the data source.</li>  
    			<li>Convert Tailwind utility classes to their valid vanilla CSS equivalent and store styles in the CSS file.</li>
				<li>If no style tags are included, insert a link to the CSS file in the HTML head.</li>
				<li>Ensure the HTML file includes a link to the generated CSS file and a script reference to the JavaScript file.</li>
				<li>The HTML file must contain the following references inside the head tag:
					<ul>
						<li><code>&lt;link rel="stylesheet" href="styles.css"&gt;</code></li>
						<li><code>&lt;script src="script.js" defer&gt;&lt;/script&gt;</code></li>
					</ul>
				</li>
				<li>This is an application component; do not generate additional explanatory text.</li>
				<li>Do not double-escape CSS selectors like: <code>.left-\\[50\\%\\]</code>. This is invalid CSS.</li>
				<li>You must escape once correctly: <code>.left-\[50\%\]</code>.</li>
				<li>Ensure transform-based animations do not interfere with other transformations. You may need to couple transformations to account for styles outside keyframes.</li>  
    		</ul>  

	</Instructions>
	`
}
