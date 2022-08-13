package repository

import (
	"fmt"

	articleEntities "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/entities"
)

// GetAllArticles returns articles to the usecases layer.
func (repository *articleRepository) GetAllArticles() (articleEntities.Articles, error) {
	// Go to DB for articles.
	articlesResponse, err := repository.database.GetAllArticlesFromDB()
	// If there is an error in the DB, return it.
	if err != nil {
		return nil, err
	}
	// Success, return articles.
	return articlesResponse, nil

}

// PostArticleIntoDB inserts articles into DB.
func (repository *articleRepository) PostArticle(article articleEntities.Article) error {
	// Send article for insert to database layer.
	err := repository.database.PostArticleIntoDB(article)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
