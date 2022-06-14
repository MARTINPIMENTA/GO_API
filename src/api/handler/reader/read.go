package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	article "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/delivery"
	"github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/home"
)

// HandleRequests function to route endpoints to their handlers.
func HandleReadRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	// home endpoint
	myRouter.HandleFunc("/", home.HomePage)
	// Get endpoint for articles
	myRouter.HandleFunc("/articles", article.GetAllArticles).Methods("GET")
	// Post endpoint for articles
	myRouter.HandleFunc("/articles", article.PostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
