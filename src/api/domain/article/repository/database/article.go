package database

import (
	"fmt"

	articleEntities "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/entities"
)

// Database mock
var dummyDB articleEntities.Articles

// GetAllArticlesFromDB gets articles from DB.
func GetAllArticlesFromDB() (articleEntities.Articles, error) {
	articlesResponse := dummyDB
	if len(articlesResponse) == 0 {
		return nil, fmt.Errorf("error getting articles from DB")
	}

	return articlesResponse, nil

}
