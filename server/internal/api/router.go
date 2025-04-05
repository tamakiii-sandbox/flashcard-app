package api

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	Swagger "github.com/swaggo/echo-swagger"
	openapi "github.com/tamakiii/flashcard-app/server/internal/api/openapi"
	"github.com/tamakiii/flashcard-app/server/internal/db"
)

func NewHandler(db *db.DB) http.Handler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", Swagger.WrapHandler)

	// Create API service with our DB
	apiService := NewFlashcardAPIService(db)

	// Create router with our custom API implementation
	apiController := openapi.NewDefaultAPIController(apiService)
	router := openapi.NewRouter(apiController)

	// Home route
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Flashcard API Server",
		})
	})

	// Mount the OpenAPI routes at root path
	e.Any("/*", echo.WrapHandler(router))

	return e
}

// FlashcardAPIService implements the DefaultAPIServicer interface from OpenAPI generator
type FlashcardAPIService struct {
	db *db.DB
}

// NewFlashcardAPIService creates an API service with database connection
func NewFlashcardAPIService(db *db.DB) openapi.DefaultAPIServicer {
	return &FlashcardAPIService{db: db}
}

// GetFlashcards implements the GET /api/flashcards endpoint
func (s *FlashcardAPIService) GetFlashcards(ctx context.Context, category string) (openapi.ImplResponse, error) {
	var rows *sql.Rows
	var err error

	if category != "" {
		rows, err = s.db.Query("SELECT id, front, back, category, created_at, updated_at FROM flashcards WHERE category = ? ORDER BY id DESC", category)
	} else {
		rows, err = s.db.Query("SELECT id, front, back, category, created_at, updated_at FROM flashcards ORDER BY id DESC")
	}

	if err != nil {
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{
			Error: "Failed to fetch flashcards",
		}), nil
	}
	defer rows.Close()

	// Use the generated Flashcard type from OpenAPI
	flashcards := []openapi.Flashcard{}
	for rows.Next() {
		var f openapi.Flashcard
		var createdAt, updatedAt time.Time
		var categoryValue sql.NullString

		if err := rows.Scan(&f.Id, &f.Front, &f.Back, &categoryValue, &createdAt, &updatedAt); err != nil {
			return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{
				Error: "Failed to scan flashcard",
			}), nil
		}

		if categoryValue.Valid {
			f.Category = categoryValue.String
		}
		f.CreatedAt = createdAt
		f.UpdatedAt = updatedAt
		flashcards = append(flashcards, f)
	}

	return openapi.Response(http.StatusOK, flashcards), nil
}

// CreateFlashcard - Create a new flashcard
func (s *FlashcardAPIService) CreateFlashcard(ctx context.Context, request openapi.FlashcardCreateRequest) (openapi.ImplResponse, error) {
	// For now, return not implemented
	return openapi.Response(http.StatusNotImplemented, nil), nil
}

// GetFlashcard - Get a specific flashcard
func (s *FlashcardAPIService) GetFlashcard(ctx context.Context, id int64) (openapi.ImplResponse, error) {
	// For now, return not implemented
	return openapi.Response(http.StatusNotImplemented, nil), nil
}

// UpdateFlashcard - Update a flashcard
func (s *FlashcardAPIService) UpdateFlashcard(ctx context.Context, id int64, request openapi.FlashcardUpdateRequest) (openapi.ImplResponse, error) {
	// For now, return not implemented
	return openapi.Response(http.StatusNotImplemented, nil), nil
}

// DeleteFlashcard - Delete a flashcard
func (s *FlashcardAPIService) DeleteFlashcard(ctx context.Context, id int64) (openapi.ImplResponse, error) {
	// For now, return not implemented
	return openapi.Response(http.StatusNotImplemented, nil), nil
}
