# Flashcard App

A cross-platform flashcard application built with React Native, Expo, and TypeScript.

## Overview

This flashcard app allows users to create, study, and manage flashcards on mobile devices (iOS and Android) and web browsers. It's designed to be a simple yet effective tool for learning and memorization.

## Technologies

- **React Native**: Core framework for mobile app development
- **Expo**: Simplifies React Native development and deployment
- **TypeScript**: Adds static typing to improve code quality and developer experience
- **Jest**: Testing framework for ensuring code reliability
- **React Testing Library**: For component testing

## Project Structure

```
flashcard-app/
├── assets/               # Contains app icons and images
├── src/                  # Source code
│   └── components/       # Reusable UI components
│       ├── Card.tsx      # Flashcard component
│       └── Card.test.tsx # Tests for Card component
├── App.tsx               # Main application component
├── app.json              # Expo configuration
├── index.ts              # Entry point
├── tsconfig.json         # TypeScript configuration
└── package.json          # Dependencies and scripts
```

## Getting Started

### Prerequisites

- Node.js (v14 or newer)
- npm or yarn
- Expo CLI
- iOS Simulator or Android Emulator (optional, for mobile development)

### Installation

1. Clone the repository:
   ```
   git clone [repository-url]
   cd flashcard-app
   ```

2. Install dependencies:
   ```
   npm install
   ```

### Running the App

#### Mobile Development

- **iOS**: 
  ```
  npm run ios
  ```

- **Android**: 
  ```
  npm run android
  ```

- **Development Server**: 
  ```
  npm start
  ```

#### Web Development

```
npm run web
```

### Testing

```
npm test
```

Run tests in watch mode:
```
npm run test:watch
```

## Development

### Creating Components

Components should be created in the `src/components` directory with corresponding test files. For example:

- `ComponentName.tsx` - The component
- `ComponentName.test.tsx` - Tests for the component

### Component Example

Here's a basic example of the Card component:

```tsx
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
```

## Best Practices

1. **TypeScript**: Always define proper interfaces for component props
2. **Testing**: Write tests for all components and functionality
3. **Components**: Create small, reusable components with a single responsibility
4. **Code Style**: Follow the established code style patterns in the project

## Contributing

1. Create a new branch for your feature or bugfix
2. Make your changes
3. Add or update tests as needed
4. Ensure all tests pass
5. Submit a pull request

## License

[Add your license information here]