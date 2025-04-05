package api

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/tamakiii/flashcard-app/server/internal/db"
	"github.com/tamakiii/flashcard-app/server/internal/models"
)

// ServerImpl implements the OpenAPI ServerInterface
type ServerImpl struct {
	db *db.DB
}

// NewServerImpl creates a new ServerImpl instance
func NewServerImpl(db *db.DB) *ServerImpl {
	return &ServerImpl{db: db}
}

// GetFlashcards handles GET /api/flashcards
func (s *ServerImpl) GetFlashcards(ctx context.Context, request GetFlashcardsRequestObject) (GetFlashcardsResponseObject, error) {
	var rows *sql.Rows
	var err error

	if request.Params.Category != nil {
		rows, err = s.db.Query("SELECT id, front, back, category, created_at, updated_at FROM flashcards WHERE category = ? ORDER BY id DESC", *request.Params.Category)
	} else {
		rows, err = s.db.Query("SELECT id, front, back, category, created_at, updated_at FROM flashcards ORDER BY id DESC")
	}

	if err != nil {
		return GetFlashcards500JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "Failed to fetch flashcards",
			},
		}, nil
	}
	defer rows.Close()

	flashcards := []Flashcard{}
	for rows.Next() {
		var f Flashcard
		if err := rows.Scan(&f.Id, &f.Front, &f.Back, &f.Category, &f.CreatedAt, &f.UpdatedAt); err != nil {
			return GetFlashcards500JSONResponse{
				ErrorResponse: ErrorResponse{
					Error: "Failed to scan flashcard",
				},
			}, nil
		}
		flashcards = append(flashcards, f)
	}

	return GetFlashcards200JSONResponse(flashcards), nil
}

// GetFlashcard handles GET /api/flashcards/{id}
func (s *ServerImpl) GetFlashcard(ctx context.Context, request GetFlashcardRequestObject) (GetFlashcardResponseObject, error) {
	var f Flashcard
	err := s.db.QueryRow("SELECT id, front, back, category, created_at, updated_at FROM flashcards WHERE id = ?", request.Id).
		Scan(&f.Id, &f.Front, &f.Back, &f.Category, &f.CreatedAt, &f.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return GetFlashcard404JSONResponse{
				ErrorResponse: ErrorResponse{
					Error: "Flashcard not found",
				},
			}, nil
		}
		return GetFlashcard500JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "Failed to fetch flashcard",
			},
		}, nil
	}

	return GetFlashcard200JSONResponse(f), nil
}

// CreateFlashcard handles POST /api/flashcards
func (s *ServerImpl) CreateFlashcard(ctx context.Context, request CreateFlashcardRequestObject) (CreateFlashcardResponseObject, error) {
	req := request.Body

	if req.Front == "" || req.Back == "" {
		return CreateFlashcard400JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "Front and back fields are required",
			},
		}, nil
	}

	result, err := s.db.Exec("INSERT INTO flashcards (front, back, category) VALUES (?, ?, ?)",
		req.Front, req.Back, req.Category)
	if err != nil {
		return CreateFlashcard500JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "Failed to create flashcard",
			},
		}, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return CreateFlashcard500JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "Failed to get flashcard ID",
			},
		}, nil
	}

	var f Flashcard
	err = s.db.QueryRow("SELECT id, front, back, category, created_at, updated_at FROM flashcards WHERE id = ?", id).
		Scan(&f.Id, &f.Front, &f.Back, &f.Category, &f.CreatedAt, &f.UpdatedAt)
	if err != nil {
		return CreateFlashcard500JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "Failed to fetch created flashcard",
			},
		}, nil
	}

	return CreateFlashcard201JSONResponse(f), nil
}

// UpdateFlashcard handles PUT /api/flashcards/{id}
func (s *ServerImpl) UpdateFlashcard(ctx context.Context, request UpdateFlashcardRequestObject) (UpdateFlashcardResponseObject, error) {
	id := request.Id
	req := request.Body

	// Check if flashcard exists
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM flashcards WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return UpdateFlashcard500JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "Failed to check flashcard existence",
			},
		}, nil
	}

	if !exists {
		return UpdateFlashcard404JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "Flashcard not found",
			},
		}, nil
	}

	// Build update query
	queryParts := []string{}
	args := []interface{}{}

	if req.Front != nil {
		queryParts = append(queryParts, "front = ?")
		args = append(args, *req.Front)
	}

	if req.Back != nil {
		queryParts = append(queryParts, "back = ?")
		args = append(args, *req.Back)
	}

	if req.Category != nil {
		queryParts = append(queryParts, "category = ?")
		args = append(args, *req.Category)
	}

	if len(queryParts) == 0 {
		return UpdateFlashcard400JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "No fields to update",
			},
		}, nil
	}

	// Execute update
	query := "UPDATE flashcards SET " + joinQueryParts(queryParts) + ", updated_at = CURRENT_TIMESTAMP WHERE id = ?"
	args = append(args, id)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		return UpdateFlashcard500JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "Failed to update flashcard",
			},
		}, nil
	}

	// Fetch updated flashcard
	var f Flashcard
	err = s.db.QueryRow("SELECT id, front, back, category, created_at, updated_at FROM flashcards WHERE id = ?", id).
		Scan(&f.Id, &f.Front, &f.Back, &f.Category, &f.CreatedAt, &f.UpdatedAt)
	if err != nil {
		return UpdateFlashcard500JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "Failed to fetch updated flashcard",
			},
		}, nil
	}

	return UpdateFlashcard200JSONResponse(f), nil
}

// DeleteFlashcard handles DELETE /api/flashcards/{id}
func (s *ServerImpl) DeleteFlashcard(ctx context.Context, request DeleteFlashcardRequestObject) (DeleteFlashcardResponseObject, error) {
	id := request.Id

	// Check if flashcard exists
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM flashcards WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return DeleteFlashcard500JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "Failed to check flashcard existence",
			},
		}, nil
	}

	if !exists {
		return DeleteFlashcard404JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "Flashcard not found",
			},
		}, nil
	}

	// Delete flashcard
	_, err = s.db.Exec("DELETE FROM flashcards WHERE id = ?", id)
	if err != nil {
		return DeleteFlashcard500JSONResponse{
			ErrorResponse: ErrorResponse{
				Error: "Failed to delete flashcard",
			},
		}, nil
	}

	return DeleteFlashcard204Response{}, nil
}

// Helper function to join query parts with commas
func joinQueryParts(parts []string) string {
	if len(parts) == 0 {
		return ""
	}

	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result += ", " + parts[i]
	}
	return result
}