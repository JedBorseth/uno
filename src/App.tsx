import { useEffect } from "react";

function App() {
  useEffect(() => {
    const socket = new WebSocket(`wss://${window.location.host}/api/ws`);
    console.log("Connecting to WebSocket...", socket);

    socket.onopen = function () {
      console.log("Connection established");
      socket.send("Hello from the client!");
    };

    socket.onmessage = function (event) {
      console.log("Message received:", event.data);
    };

    socket.onerror = function (error) {
      console.error("WebSocket error:", error, socket);
    };
  }, []);

  return (
    <div className="flex items-center justify-center h-screen bg-gray-100">
      <h1 className="text-4xl">Hello World</h1>
    </div>
  );
}

export default App;
