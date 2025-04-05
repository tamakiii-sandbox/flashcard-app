import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

describe('App component', () => {
  test('renders the app header', () => {
    render(<App />);
    expect(screen.getByText('Flashcard App')).toBeInTheDocument();
  });

  test('renders a card component', () => {
    render(<App />);
    expect(screen.getByTestId('flashcard')).toBeInTheDocument();
  });
});
