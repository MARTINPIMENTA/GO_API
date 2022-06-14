package article

import (
	"encoding/json"
	"fmt"
	"net/http"
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

// AllArticles returns the array of articles.
func AllArticles(w http.ResponseWriter, r *http.Request) {
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
func PostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This also works, POST method")
}
