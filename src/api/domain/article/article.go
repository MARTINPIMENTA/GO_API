package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

var articles Articles

// AllArticles returns the array of articles.
func AllArticles(w http.ResponseWriter, r *http.Request) {
	// simulation to DB Get request
	articlesResponse := articles

	// Empty DB error case
	if len(articlesResponse) == 0 {
		fmt.Println("Endpoint Hit with error, no articles")
		fmt.Fprintf(w, "No articles to return from DB")
		return
	}

	// Success
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

// testPostArticles only returns a message, dummy func.
func PostArticles(w http.ResponseWriter, r *http.Request) {
	article := new(Article)

	// Body error case
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error formatting request")
		return
	}

	// Unmarshal body error case
	err = json.Unmarshal(body, article)
	if err != nil {
		fmt.Fprintf(w, "Body unmarshall error")
		return
	}

	// Success
	articles = append(articles, *article)
	fmt.Fprintf(w, "Article Posted!")
}
