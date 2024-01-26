import './App.css';

function App() {
  const divStyle = {
    margin: '40px',
    border: '5px solid pink'
  };
  return (
    <div className="App">
      <header className="App-header">
        <button style={divStyle} onClick= getForm('A')> survey A</button>
        <button style={divStyle}> survey B</button>
        <button style={divStyle}> survey C</button>
      </header>
    </div>
  );
}

export default App;
