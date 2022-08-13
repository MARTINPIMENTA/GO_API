package repository

import (
	articleEntities "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/entities"
	"github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/repository/database"
	"github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/infrastructure/dependencies"
)

type ArticleRepository interface {
	GetAllArticles() (articleEntities.Articles, error)

	PostArticle(article articleEntities.Article) error
}

type articleRepository struct {
	database database.ArticleDatabase
}

func NewRepository(container *dependencies.Container) ArticleRepository {
	database := database.NewArticleDatabaseRepository(container)
	return &articleRepository{
		database: database,
	}
}
