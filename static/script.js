// Initiate WebSocket connection
const ws = new WebSocket("ws://localhost:8080/ws");

// WebSocket event listeners
ws.onopen = () => {
  console.log("Connected");

  // Send message to WebSocket server
  ws.onmessage = (message) => {
    console.log(message.data);
  };

  // Close WebSocket connection
  ws.onclose = () => {
    console.log("Disconnected");
  };
};

// DOM manipulation
document.addEventListener("DOMContentLoaded", function () {
  const title = document.getElementById("title");
  const changeTextBtn = document.getElementById("changeTextBtn");

  changeTextBtn.addEventListener("click", function () {
    title.textContent = "Golang Rocks!";
  });
});
