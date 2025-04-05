package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/tamakiii/flashcard-app/server/internal/db"
	"github.com/tamakiii/flashcard-app/server/internal/models"
)

// Handler handles HTTP requests
type Handler struct {
	db *db.DB
}

// NewHandler creates a new Handler instance
func NewHandler(db *db.DB) *Handler {
	return &Handler{db: db}
}

// respondWithJSON sends a JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError sends an error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// GetFlashcards handles GET /api/flashcards
func (h *Handler) GetFlashcards(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	var rows *sql.Rows
	var err error

	if category != "" {
		rows, err = h.db.Query("SELECT id, front, back, category, created_at, updated_at FROM flashcards WHERE category = ? ORDER BY id DESC", category)
	} else {
		rows, err = h.db.Query("SELECT id, front, back, category, created_at, updated_at FROM flashcards ORDER BY id DESC")
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch flashcards")
		return
	}
	defer rows.Close()

	flashcards := []models.Flashcard{}
	for rows.Next() {
		var f models.Flashcard
		if err := rows.Scan(&f.ID, &f.Front, &f.Back, &f.Category, &f.CreatedAt, &f.UpdatedAt); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to scan flashcard")
			return
		}
		flashcards = append(flashcards, f)
	}

	respondWithJSON(w, http.StatusOK, flashcards)
}

// GetFlashcard handles GET /api/flashcards/:id
func (h *Handler) GetFlashcard(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(getIDFromPath(r.URL.Path), 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid flashcard ID")
		return
	}

	var f models.Flashcard
	err = h.db.QueryRow("SELECT id, front, back, category, created_at, updated_at FROM flashcards WHERE id = ?", id).
		Scan(&f.ID, &f.Front, &f.Back, &f.Category, &f.CreatedAt, &f.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "Flashcard not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Failed to fetch flashcard")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, f)
}

// CreateFlashcard handles POST /api/flashcards
func (h *Handler) CreateFlashcard(w http.ResponseWriter, r *http.Request) {
	var req models.FlashcardCreateRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if req.Front == "" || req.Back == "" {
		respondWithError(w, http.StatusBadRequest, "Front and back fields are required")
		return
	}

	result, err := h.db.Exec("INSERT INTO flashcards (front, back, category) VALUES (?, ?, ?)",
		req.Front, req.Back, req.Category)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create flashcard")
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get flashcard ID")
		return
	}

	var f models.Flashcard
	err = h.db.QueryRow("SELECT id, front, back, category, created_at, updated_at FROM flashcards WHERE id = ?", id).
		Scan(&f.ID, &f.Front, &f.Back, &f.Category, &f.CreatedAt, &f.UpdatedAt)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch created flashcard")
		return
	}

	respondWithJSON(w, http.StatusCreated, f)
}

// UpdateFlashcard handles PUT /api/flashcards/:id
func (h *Handler) UpdateFlashcard(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(getIDFromPath(r.URL.Path), 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid flashcard ID")
		return
	}

	var req models.FlashcardUpdateRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Check if flashcard exists
	var exists bool
	err = h.db.QueryRow("SELECT EXISTS(SELECT 1 FROM flashcards WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to check flashcard existence")
		return
	}

	if !exists {
		respondWithError(w, http.StatusNotFound, "Flashcard not found")
		return
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
		respondWithError(w, http.StatusBadRequest, "No fields to update")
		return
	}

	queryParts = append(queryParts, "updated_at = CURRENT_TIMESTAMP")
	query := "UPDATE flashcards SET " + strings.Join(queryParts, ", ") + " WHERE id = ?"
	args = append(args, id)

	_, err = h.db.Exec(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update flashcard")
		return
	}

	var f models.Flashcard
	err = h.db.QueryRow("SELECT id, front, back, category, created_at, updated_at FROM flashcards WHERE id = ?", id).
		Scan(&f.ID, &f.Front, &f.Back, &f.Category, &f.CreatedAt, &f.UpdatedAt)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch updated flashcard")
		return
	}

	respondWithJSON(w, http.StatusOK, f)
}

// DeleteFlashcard handles DELETE /api/flashcards/:id
func (h *Handler) DeleteFlashcard(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(getIDFromPath(r.URL.Path), 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid flashcard ID")
		return
	}

	// Check if flashcard exists
	var exists bool
	err = h.db.QueryRow("SELECT EXISTS(SELECT 1 FROM flashcards WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to check flashcard existence")
		return
	}

	if !exists {
		respondWithError(w, http.StatusNotFound, "Flashcard not found")
		return
	}

	_, err = h.db.Exec("DELETE FROM flashcards WHERE id = ?", id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to delete flashcard")
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}

// getIDFromPath extracts the ID from the URL path
func getIDFromPath(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}
