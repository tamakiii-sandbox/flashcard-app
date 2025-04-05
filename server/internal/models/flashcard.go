package models

import (
	"time"
)

// Flashcard represents a single flashcard
type Flashcard struct {
	ID        int64     `json:"id"`
	Front     string    `json:"front"`
	Back      string    `json:"back"`
	Category  string    `json:"category,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FlashcardCreateRequest represents a request to create a new flashcard
type FlashcardCreateRequest struct {
	Front    string `json:"front"`
	Back     string `json:"back"`
	Category string `json:"category,omitempty"`
}

// FlashcardUpdateRequest represents a request to update an existing flashcard
type FlashcardUpdateRequest struct {
	Front    *string `json:"front,omitempty"`
	Back     *string `json:"back,omitempty"`
	Category *string `json:"category,omitempty"`
}
