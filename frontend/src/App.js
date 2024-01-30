import './App.css';
import DynamicForm from './component/dynamic_form.jsx';

function App() {
  const divStyle = {
    margin: '40px',
    border: '5px solid pink'
  };
  return (
    <div className="App">
      <header className="App-header">
        <DynamicForm />
      </header>
    </div>
  );
}

export default App;
