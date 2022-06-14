package usecases

import (
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
