package handler

import (
	"github.com/gorilla/mux"

	article "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/delivery"
	"github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/home"
)

// HandleRequests function to route endpoints to their handlers.
func HandleReadRequests(router *mux.Router) {

	// home endpoint
	router.HandleFunc("/", home.HomePage)
	// Get endpoint for articles
	router.HandleFunc("/articles", article.GetAllArticles).Methods("GET")
}
