package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	articleEntities "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/entities"
	articleUsecase "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/usecases"
)

// GetAllArticles returns the array of articles.
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	// Verify request is not empty
	err := verifyValidRequest(r)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	// Go to usecase directly because there is no specific article.
	articlesResponse, err := articleUsecase.GetAllArticles()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
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

// PostArticles sends articles to DB.
func PostArticles(w http.ResponseWriter, r *http.Request) {
	// Verify request is not empty
	err := verifyValidRequest(r)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	// Body error case
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error formatting request")
		return
	}

	article := new(articleEntities.Article)

	// Unmarshal body error case
	err = json.Unmarshal(body, article)
	if err != nil {
		fmt.Fprintf(w, "Body unmarshall error")
		return
	}

	// Go to usecase for article post.
	err = articleUsecase.PostArticle(*article)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	// Success
	fmt.Fprintf(w, "Article Posted!")
}

// getRequestHeaders gets request params from URL.
/*func getRequestParams(r *http.Request) error {
	// Get the value

}*/

func verifyValidRequest(r *http.Request) error {
	if r == nil {
		return fmt.Errorf("Request error")
	}
	return nil
}
