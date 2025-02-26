import { useState, useEffect } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [message, setMessage] = useState('')
  const [ws, setWs] = useState<WebSocket | null>(null);

  useEffect(() => {
    const userId = 'zach';
    const wsUrl = 'ws://localhost:5000/ws/' + userId;
    //const wsUrl = 'ws://localhost:5000/ws';

    const socket = new WebSocket(wsUrl);

    setWs(socket);

    socket.onmessage = (event) => {
      setMessage(event.data);
    };

    socket.onopen = () => {
      console.log('WebSocket connection opened');
    };

    socket.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    socket.onclose = () => {
      console.log('WebSocket connection closed');
    };
    return () => {
      socket.close();
    };
  }, []);

  const handleSendMessage = () => {
    if (ws) {
      ws.send('Hello from React!');
    }
  };


  return (
    <>
      <div>
        <img src={viteLogo} className="logo" alt="Vite logo" />    
        <img src={reactLogo} className="logo react" alt="React logo" />
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <h1>WebSocket Demo</h1>
        <p>Message: {message}</p>
        <button onClick={handleSendMessage}>Send Message</button>
      </div>
    </>
  )
}

export default App
