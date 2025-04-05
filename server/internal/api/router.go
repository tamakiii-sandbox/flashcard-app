package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tamakiii/flashcard-app/server/internal/db"
)

func NewHandler(db *db.DB) http.Handler {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return e
}

// // Router handles HTTP routing with OpenAPI validation
// type Router struct {
// 	handler ServerInterface
// 	// swagger  *Swagger
// 	basePath string
// }

// // NewRouter creates a new Router instance
// func NewRouter(db *db.DB) *Router {
// 	swagger, err := GetSwagger()
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Clear the ServerURLs to avoid validation issues
// 	swagger.Servers = nil

// 	return &Router{
// 		handler:  NewServerImpl(db),
// 		swagger:  swagger,
// 		basePath: "",
// 	}
// }

// // ServeHTTP implements the http.Handler interface
// func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	// Set up CORS headers
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 	// Handle preflight requests
// 	if req.Method == http.MethodOptions {
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}

// 	// Create Chi router
// 	router := chi.NewRouter()

// 	// Add middleware for OpenAPI validation
// 	router.Use(middleware.OapiRequestValidator(r.swagger))

// 	// Register handler
// 	handler := NewStrictHandler(r.handler, nil)
// 	HandlerFromMux(handler, router)

// 	// Serve the request
// 	router.ServeHTTP(w, req)
// }
