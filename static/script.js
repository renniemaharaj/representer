
			// Create a WebSocket connection
			const ws = new WebSocket("ws://localhost:8080/ws");

			ws.onopen = function(event) {
				console.log("WebSocket connection opened.");

				// Send a message to the server
				ws.send("Hello, server!");

				// Receive messages from the server
				ws.onmessage = function(event) {
					document.querySelector("#title").innerHTML = event.data;
				};
			};
		