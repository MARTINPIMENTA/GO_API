package article

import (
	"encoding/json"
	"fmt"
	"net/http"

	articleUsecase "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/usecases"
)

// AllArticles returns the array of articles.
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	// Verify request is not empty
	if r == nil {
		fmt.Fprintf(w, "Empty request error")
		return
	}

	// Go to usecase directly because there is no specific article.
	articlesResponse, err := articleUsecase.GetAllArticles()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// Empty DB error case
	if len(articlesResponse) == 0 {
		fmt.Println("Endpoint Hit with error, no articles")
		fmt.Fprintf(w, "No articles to return from DB")
		return
	}

	// Success
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articlesResponse)
}

// testPostArticles only returns a message, dummy func.
// func PostArticles(w http.ResponseWriter, r *http.Request) {
// 	article := new(articleEntities.Article)

// 	// Body error case
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Error formatting request")
// 		return
// 	}

// 	// Unmarshal body error case
// 	err = json.Unmarshal(body, article)
// 	if err != nil {
// 		fmt.Fprintf(w, "Body unmarshall error")
// 		return
// 	}

// 	// Success
// 	articles = append(articles, *article)
// 	fmt.Fprintf(w, "Article Posted!")
// }

/* func getRequestHeaders(r *http.Request) error {
	// Get

} */
