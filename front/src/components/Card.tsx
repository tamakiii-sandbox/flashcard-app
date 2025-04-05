import React from 'react';

interface CardProps {
  question: string;
  answer: string;
}

export const Card: React.FC<CardProps> = ({ question, answer }) => {
  const [showAnswer, setShowAnswer] = React.useState(false);

  return (
    <div className="card" data-testid="flashcard">
      <div className="card-content">
        <div className="question">{question}</div>
        {showAnswer && <div className="answer">{answer}</div>}
      </div>
      <button 
        onClick={() => setShowAnswer(!showAnswer)}
        data-testid="toggle-button"
      >
        {showAnswer ? 'Hide Answer' : 'Show Answer'}
      </button>
    </div>
  );
};
