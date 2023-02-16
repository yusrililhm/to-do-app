import './assets/css/index.css';
import Header from './components/header';
import ToDoList from './content/todo';

function App() {
  return (
    <section className="font-Poppins">
      <Header />
      <ToDoList />
    </section>
  );
}

export default App;
