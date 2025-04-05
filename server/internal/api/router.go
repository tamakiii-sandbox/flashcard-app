package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	Swagger "github.com/swaggo/echo-swagger"
	"github.com/tamakiii/flashcard-app/server/internal/db"
)

func NewHandler(db *db.DB) http.Handler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", Swagger.WrapHandler)

	// Register the API endpoints according to our OpenAPI specification
	api := e.Group("/api")
	{
		api.GET("/flashcards", func(c echo.Context) error {
			return getFlashcards(c, db)
		})
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello",
		})
	})

	return e
}

// getFlashcards implements the GET /api/flashcards endpoint
func getFlashcards(c echo.Context, db *db.DB) error {
	var rows *sql.Rows
	var err error

	// Query the database based on category parameter
	rows, err = db.Query("SELECT id, front, back, category, created_at, updated_at FROM flashcards ORDER BY id DESC")

	if err != nil {
		// Using the generated ErrorResponse type from api.gen.go
		errResp := ErrorResponse{
			Error: "Failed to fetch flashcards",
		}
		return c.JSON(http.StatusInternalServerError, errResp)
	}
	defer rows.Close()

	// We're using the generated Flashcard type from api.gen.go
	results := []Flashcard{}
	for rows.Next() {
		var f Flashcard

		if err := rows.Scan(&f.Id, &f.Front, &f.Back, &f.CreatedAt, &f.UpdatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error: "Failed to scan flashcard",
			})
		}

		results = append(results, f)
	}

	return c.JSON(http.StatusOK, results)
}
