package database

import (
	articleEntities "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/entities"
	"github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/infrastructure/dependencies"
)

type ArticleDatabase interface {
	// GetAllArticlesFromDB method to get articles from DB.
	GetAllArticlesFromDB() (articleEntities.Articles, error)
	// PostArticleIntoDB method to insert articles into DB.
	PostArticleIntoDB(article articleEntities.Article) error
}

type articleDatabase struct {
	database string
}

func NewArticleDatabaseRepository(container *dependencies.Container) ArticleDatabase {
	return &articleDatabase{
		database: container.Database(),
	}
}

// NewMockArticleDatabaseRepository to mock database struct.
// func NewMockArticleDatabaseRepository(mockField *string) ArticleDatabase {
// 	return &articleDatabase{
// 		mockField: "",
// 	}
// }
