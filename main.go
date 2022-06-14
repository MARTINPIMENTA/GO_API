package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	lambdaFuncs "src/api/domain/lambda/entities"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/mux"
)

// Entities for article functionalities.
type (
	Article struct {
		Title   string `json:"title"`
		Desc    string `json:"desc"`
		Content string `json:"content"`
	}

	Articles []Article
)

// allArticles returns the array of articles.
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "Test Title",
			Desc:    "Test Description",
			Content: "Hello World",
		},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

// testPostArticles only returns a message, dummy func.
func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This also works, POST method")
}

// homePage func to test if the API is up and running.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

// handleRequests function to route endpoints to their handlers.
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	// home endpoint
	myRouter.HandleFunc("/", homePage)
	// Get endpoint for articles
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	// Post endpoint for articles
	myRouter.HandleFunc("/articlespost", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

// main func.
func main() {
	handleRequests()

	// Initialization for lambda functions.
	lambda.Start(lambdaFuncs.HandlerLambda)
}
