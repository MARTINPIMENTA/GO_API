package usecases

import (
	"fmt"

	articleEntities "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/entities"
	articleRepository "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/repository"
)

func GetAllArticles() (articleEntities.Articles, error) {
	// Usecase doesn`t do much in this case, there are no parameters to check.
	// Go to repository and check for errors.
	articlesResponse, err := articleRepository.GetAllArticles()
	if err != nil {
		return nil, err
	}

	return articlesResponse, nil
}

// PostArticleIntoDB inserts articles into DB.
func PostArticle(article articleEntities.Article) error {
	// Check if the article is valid.
	if !isValidArticle(article) {
		return fmt.Errorf("invalid article, empty title")
	}

	// Send article for insert to repository layer.
	err := articleRepository.PostArticle(article)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

// Checks article for invalid cases.
func isValidArticle(article articleEntities.Article) bool {
	// If title is empty, then the article is invalid.
	return article.Title != ""
}
