import { useState } from 'react';
import './App.css';
import { CalculateForm } from './components/CalculateForm/CalculateForm';
import { Details } from './components/Details/Details';
import { Header } from './components/Header/Header';

function App() {
  const [state, setState] = useState({ data: null, error: null })

  const onResponse = (resp) => {
    if (!resp.ok) {
      return setState((s) => ({ ...s, error: resp.json }))
    }

    return setState((s) => ({ ...s, data: resp.json }))
  }

  return (
    <div className="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="w-full max-w-screen-lg">
        <Header />
        <section className="grid grid-cols-2 gap-2 place-content-stretch">
          <CalculateForm onResponse={onResponse} />
          <Details data={state.data} />
        </section>
      </div>
    </div>
  );
}

export default App;
