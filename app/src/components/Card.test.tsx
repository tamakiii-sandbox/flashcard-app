import React from 'react';
import { render, screen } from '@testing-library/react-native';
import { Card } from './Card';

describe('Card Component', () => {
  it('renders the title correctly', () => {
    render(<Card title="Test Title" />);
    
    const titleElement = screen.getByText('Test Title');
    expect(titleElement).toBeTruthy();
  });

  it('renders the content when provided', () => {
    render(<Card title="Test Title" content="Test Content" />);
    
    const contentElement = screen.getByText('Test Content');
    expect(contentElement).toBeTruthy();
  });

  it('does not render content when not provided', () => {
    render(<Card title="Test Title" />);
    
    const contentElements = screen.queryAllByText(/Content/);
    expect(contentElements.length).toBe(0);
  });

  it('has the correct testID', () => {
    render(<Card title="Test Title" />);
    
    const cardElement = screen.getByTestId('flashcard');
    expect(cardElement).toBeTruthy();
  });
});
