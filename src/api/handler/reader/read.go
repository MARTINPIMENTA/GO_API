package handler

import (
	"github.com/gorilla/mux"

	article "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/delivery"
	"github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/home"
	"github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/infrastructure/dependencies"
)

type Reader struct {
	container *dependencies.Container
}

// NewReader returns a concrete implementation of the Reader scope
func NewReader(container *dependencies.Container) *Reader {
	return &Reader{
		container: container,
	}
}

// HandleRequests function to route endpoints to their handlers.
func (reader *Reader) HandleReadRequests(router *mux.Router) {
	articleHandler := article.NewArticleHandlerHTTP(reader.container)

	// home endpoint
	router.HandleFunc("/", home.HomePage)
	// Get endpoint for articles
	router.HandleFunc("/articles", articleHandler.GetAllArticles).Methods("GET")
}
