import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import { Card } from './Card';

describe('Card component', () => {
  const mockProps = {
    question: 'What is React?',
    answer: 'A JavaScript library for building user interfaces'
  };

  test('renders the question', () => {
    render(<Card {...mockProps} />);
    expect(screen.getByText(mockProps.question)).toBeInTheDocument();
  });

  test('initially hides the answer', () => {
    render(<Card {...mockProps} />);
    expect(screen.queryByText(mockProps.answer)).not.toBeInTheDocument();
  });

  test('shows the answer when the button is clicked', () => {
    render(<Card {...mockProps} />);
    fireEvent.click(screen.getByTestId('toggle-button'));
    expect(screen.getByText(mockProps.answer)).toBeInTheDocument();
  });

  test('hides the answer when the button is clicked again', () => {
    render(<Card {...mockProps} />);
    const button = screen.getByTestId('toggle-button');
    
    // Show answer
    fireEvent.click(button);
    expect(screen.getByText(mockProps.answer)).toBeInTheDocument();
    
    // Hide answer
    fireEvent.click(button);
    expect(screen.queryByText(mockProps.answer)).not.toBeInTheDocument();
  });
});
