import React from 'react';

function App() {
  const handleLogin = () => {
    window.location.href = 'http://localhost:8080/login';
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>Connexion avec Google</h1>
        <button onClick={handleLogin}>Se connecter avec Google</button>
      </header>
    </div>
  );
}

export default App;
