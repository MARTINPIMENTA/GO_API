package writer

import (
	"log"
	"net/http"

	"github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article"
	"github.com/gorilla/mux"
)

// HandleRequests function to route endpoints to their handlers.
func HandleWriteRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	// Post endpoint for articles
	myRouter.HandleFunc("/articles", article.PostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
