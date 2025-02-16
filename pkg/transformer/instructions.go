package transformer

// GetInstructions returns the system instructions
func GetInstructions() string {
	return `
<Instructions>

    <h1>Responsive Web Page Generator</h1>  

    <h2>Request</h2>  
        <p>Generate a responsive web page document with separate HTML and CSS files, based on a marshalled document representation.</p>  

    <h2>Description</h2>  
        <p>This application generates responsive web pages from a provided JSON object. The object represents a Go struct-based document, defining the structure of the web document. Elements use Tailwind or similarly styled utility classes in their class attributes.</p>  

    <h2>Response Format</h2>
        <p>The response must be a JSON object containing two files: an HTML file and a CSS file. The structure must be as follows:</p>
        <pre>
        {
            "html": {
                "content": "HTML_CONTENT",
                "filename": "index.html"
            },
            "css": {
                "content": "CSS_CONTENT",
                "filename": "styles.css"
            }
        }
        </pre>
        <p>If the document object already contains a linked CSS file with the default filename ("styles.css"), you must generate and use an alternative filename.</p>
        <p>User-defined filenames take priority. However, if a user links a stylesheet with the default name ("styles.css"), generate your own CSS file with a unique name to avoid conflicts.</p>
        <p>Example logic:</p>
        <ul>
            <li>If no CSS file is linked or no style element is embedded in the document, generate "styles.css".</li>
            <li>If a user has already linked "styles.css", generate a different filename (e.g., "g-styles.css").</li>
            <li>Ensure the generated HTML references the correct filename for consistency.</li>
        </ul>

    <h2>Instructions</h2>  
        <ul>  
            <li>Use the attached JSON object as the data source.</li>  
            <li>Convert Tailwind utility classes to their valid vanilla CSS equivalent and store styles in the CSS file.</li>
            <li>Do not omit class conversations for elements, all elements are required to have vanilla CSS equivalents for their utility class declaration<li>
            <li>If no style tags are included, insert a link to the CSS file in the HTML head.</li>
            <li>Ensure the HTML file includes:
                <ul>
                    <li>A reference to the generated CSS file.</li>
                    <li>Any embedded styles or linked stylesheets found in the marshalled document representation.</li>
                    <li>Any user-defined scripts found in the marshalled document representation.</li>
                </ul>
            </li>
            <li>Include all JavaScript files defined in the document's scripts section.</li>
            <li>Scripts must be placed in the document's head or at the end of the body, following standard HTML practices.</li>
            <li>Do not generate additional explanatory text—this is an application component.</li>
            <li>Do not double-escape CSS selectors like: <code>.left-\\[50\\%\\]</code>. This is invalid CSS.</li>
            <li>You must escape once correctly: <code>.left-\[50\%\]</code>.</li>
            <li>Ensure transform-based animations do not interfere with other transformations. You may need to couple transformations to account for styles outside keyframes.</li>  
        </ul>  

</Instructions>
	`
}
