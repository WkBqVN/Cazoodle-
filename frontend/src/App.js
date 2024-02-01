import './App.css';
import NewForm from './component/newForm.jsx';

function App() {
  const divStyle = {
    margin: '40px',
    border: '5px solid pink'
  };
  return (
    <div className="App">
      <header className="App-header">
        <NewForm />
      </header>
    </div>
  );
}

export default App;
