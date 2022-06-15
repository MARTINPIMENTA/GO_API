package writer

import (
	article "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/delivery"
	"github.com/gorilla/mux"
)

// HandleRequests function to route endpoints to their handlers.
func HandleWriteRequests(router *mux.Router) {
	// Post endpoint for articles
	router.HandleFunc("/articles", article.PostArticles).Methods("POST")
}
