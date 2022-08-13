package usecases

import (
	articleEntities "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/entities"
	"github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/repository"
	"github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/infrastructure/dependencies"
)

type ArticleUsecase interface {
	//
	GetAllArticles() (articleEntities.Articles, error)
	//
	PostArticle(articleEntities.Article) error
}

type articleUsecase struct {
	repository repository.ArticleRepository
}

func NewArticleUsecase(container *dependencies.Container) ArticleUsecase {
	repository := repository.NewRepository(container)
	return &articleUsecase{
		repository: repository,
	}
}
