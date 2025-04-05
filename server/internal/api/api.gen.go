// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Flashcard defines model for Flashcard.
type Flashcard struct {
	Back      string    `json:"back"`
	Category  *string   `json:"category,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Front     string    `json:"front"`
	Id        int64     `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FlashcardCreateRequest defines model for FlashcardCreateRequest.
type FlashcardCreateRequest struct {
	Back     string  `json:"back"`
	Category *string `json:"category,omitempty"`
	Front    string  `json:"front"`
}

// FlashcardUpdateRequest defines model for FlashcardUpdateRequest.
type FlashcardUpdateRequest struct {
	Back     *string `json:"back,omitempty"`
	Category *string `json:"category,omitempty"`
	Front    *string `json:"front,omitempty"`
}

// BadRequest defines model for BadRequest.
type BadRequest = ErrorResponse

// InternalServerError defines model for InternalServerError.
type InternalServerError = ErrorResponse

// NotFound defines model for NotFound.
type NotFound = ErrorResponse

// GetFlashcardsParams defines parameters for GetFlashcards.
type GetFlashcardsParams struct {
	// Category Filter flashcards by category
	Category *string `form:"category,omitempty" json:"category,omitempty"`
}

// CreateFlashcardJSONRequestBody defines body for CreateFlashcard for application/json ContentType.
type CreateFlashcardJSONRequestBody = FlashcardCreateRequest

// UpdateFlashcardJSONRequestBody defines body for UpdateFlashcard for application/json ContentType.
type UpdateFlashcardJSONRequestBody = FlashcardUpdateRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all flashcards
	// (GET /api/flashcards)
	GetFlashcards(ctx echo.Context, params GetFlashcardsParams) error
	// Create a new flashcard
	// (POST /api/flashcards)
	CreateFlashcard(ctx echo.Context) error
	// Delete a flashcard
	// (DELETE /api/flashcards/{id})
	DeleteFlashcard(ctx echo.Context, id int64) error
	// Get a specific flashcard
	// (GET /api/flashcards/{id})
	GetFlashcard(ctx echo.Context, id int64) error
	// Update a flashcard
	// (PUT /api/flashcards/{id})
	UpdateFlashcard(ctx echo.Context, id int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetFlashcards converts echo context to params.
func (w *ServerInterfaceWrapper) GetFlashcards(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFlashcardsParams
	// ------------- Optional query parameter "category" -------------

	err = runtime.BindQueryParameter("form", true, false, "category", ctx.QueryParams(), &params.Category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter category: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetFlashcards(ctx, params)
	return err
}

// CreateFlashcard converts echo context to params.
func (w *ServerInterfaceWrapper) CreateFlashcard(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateFlashcard(ctx)
	return err
}

// DeleteFlashcard converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteFlashcard(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteFlashcard(ctx, id)
	return err
}

// GetFlashcard converts echo context to params.
func (w *ServerInterfaceWrapper) GetFlashcard(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetFlashcard(ctx, id)
	return err
}

// UpdateFlashcard converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateFlashcard(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateFlashcard(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/flashcards", wrapper.GetFlashcards)
	router.POST(baseURL+"/api/flashcards", wrapper.CreateFlashcard)
	router.DELETE(baseURL+"/api/flashcards/:id", wrapper.DeleteFlashcard)
	router.GET(baseURL+"/api/flashcards/:id", wrapper.GetFlashcard)
	router.PUT(baseURL+"/api/flashcards/:id", wrapper.UpdateFlashcard)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RX32/bNhD+Vwhuj5rlrNlQ6GVrl3kwUAxBimEPRTBcqJPFTiJZ8uTWCPS/DyRl64e1",
	"xMUMB3uzJN59x+87fkc/cqFroxUqcjx75Bad0cpheHgL+R1+atCRfxJaEarwE4yppACSWqUfnVb+nRMl",
	"1uB/fWux4Bn/Ju1Tp/GrS3+1Vtu7DoS3bZvwHJ2w0vhkPPOYzHagbcLXitAqqN6j3aIN0ZerZQ/OXEBn",
	"GODbhP+uaaUblV+ulDt0urECmdLEioDtF3XxPv04RfbIjdUGLcmoJe6pwy9Qmwp5FiNYjc7BBnnCaWf8",
	"a0dWqk1I74WQFnOefegS3B+W6YePKIJGqwpcKcDmx6gPIP4eg96Cle4YLOECCDfa7sbLf0O9sWDK3WyI",
	"RSDM/4LAfqFt7X/xHAi/I1njXExhdRSrx/izBGLSMSqRCTCSoGK6YCsLSuBPc0lkPspwlfTgUtGP132M",
	"VIQbDE3TmPwrq50oIHO+rz+JzI4oGCE8qdMvIWhwtC8p2hkEmPAyouTJjf8RCPpfb3yyOf9KqkLvnQhE",
	"hKhBVj6wMUZb+rmDXAhd84QrqH2O9/EjPzKbN7drVnhrAAUbqTas2BPomFQMWIVglf8wcDxfraSwrQPf",
	"7M3tmid8i9bFzFeL5WLpAbVBBUbyjL9aLBeveMINUBlUSMHItEf0rzYYdjW1RLISt+h8QdKRZw6qalAs",
	"Dzg2lLfOgz60Gn41YKFGQut49mGafyUrQjvc+8OOHXT3tPOMf2owPHSUDj73xj8V8T4Zj9nvl8uvGiSS",
	"sHbPTZTelPumAWthNzdd3nX0DahrE/5DrGsO5VB/Ojegw2Rq6hr8+fCkT3VpE260m9E0OpNXVOHnPoJ9",
	"llSGU2Ks3socc5YDwZG+MbzfevQJdPRW57uzDet/8dF27EtkG2yPlL46fxVzgvYnsBsQzDVCoHNFU1U7",
	"z//1KeIOLoDn64fI2lTisGhy9NNHmbexSSokPG6Xm/Det0vfKg87Jsmx9c1Rd8TV4+4YiXN9jNATGWuY",
	"I/L6eVYOF8bz0Rh3M9y6T/6sU55C1NAm+X80qzO0cLFn7qWoDg7GnEEhCynGhD85QtY33lW9bxUDPsPg",
	"8OOunxty71W9dwwnyLO3Sz9TTDMjfbzxOAaK4RfpaDTNT7XVmORitjq+pZ1kqxfvye6qfQ5bfamujjSP",
	"DSSsCBFz7fxOC6hYjlustKlRUffH2P/1sP66WRKZLE0rv67UjrLXy9dL3t63/wQAAP//LHPFl2cQAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
