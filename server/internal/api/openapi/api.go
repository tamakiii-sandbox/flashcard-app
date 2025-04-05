// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Flashcard API
 *
 * API for managing flashcards in a learning application
 *
 * API version: 1.0.0
 * Contact: support@example.com
 */

package api

import (
	"context"
	"net/http"
)



// DefaultAPIRouter defines the required methods for binding the api requests to a responses for the DefaultAPI
// The DefaultAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultAPIServicer to perform the required actions, then write the service results to the http response.
type DefaultAPIRouter interface { 
	GetFlashcards(http.ResponseWriter, *http.Request)
	CreateFlashcard(http.ResponseWriter, *http.Request)
	GetFlashcard(http.ResponseWriter, *http.Request)
	UpdateFlashcard(http.ResponseWriter, *http.Request)
	DeleteFlashcard(http.ResponseWriter, *http.Request)
}


// DefaultAPIServicer defines the api actions for the DefaultAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultAPIServicer interface { 
	GetFlashcards(context.Context, string) (ImplResponse, error)
	CreateFlashcard(context.Context, FlashcardCreateRequest) (ImplResponse, error)
	GetFlashcard(context.Context, int64) (ImplResponse, error)
	UpdateFlashcard(context.Context, int64, FlashcardUpdateRequest) (ImplResponse, error)
	DeleteFlashcard(context.Context, int64) (ImplResponse, error)
}
