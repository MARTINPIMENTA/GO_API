package writer

import (
	article "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/delivery"
	"github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/infrastructure/dependencies"
	"github.com/gorilla/mux"
)

type Writer struct {
	container *dependencies.Container
}

// NewWriter returns a concrete implementation of the Writer scope
func NewWriter(container *dependencies.Container) *Writer {
	return &Writer{
		container: container,
	}
}

// HandleRequests function to route endpoints to their handlers.
func (writer *Writer) HandleWriteRequests(router *mux.Router) {
	articleHandler := article.NewArticleHandlerHTTP(writer.container)
	// Post endpoint for articles
	router.HandleFunc("/articles", articleHandler.PostArticles).Methods("POST")
}
