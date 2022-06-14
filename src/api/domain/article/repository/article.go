package repository

import (
	articleEntities "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/entities"
	articleDatabase "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/repository/database"
)

// GetAllArticles returns articles to the usecases layer.
func GetAllArticles() (articleEntities.Articles, error) {
	// Go to DB for articles.
	articlesResponse, err := articleDatabase.GetAllArticlesFromDB()
	// If there is an error in the DB, return it.
	if err != nil {
		return nil, err
	}
	// Success, return articles.
	return articlesResponse, nil

}
