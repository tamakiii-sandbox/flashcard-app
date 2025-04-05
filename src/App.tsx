import React from 'react';
import { Card } from './components/Card';

const App: React.FC = () => {
  return (
    <div className="app">
      <header>
        <h1>Flashcard App</h1>
      </header>
      <main>
        <Card 
          question="What is TypeScript?" 
          answer="TypeScript is a strongly typed programming language that builds on JavaScript, giving you better tooling at any scale."
        />
      </main>
    </div>
  );
};

export default App;