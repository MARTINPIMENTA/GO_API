package article

import (
	"net/http"

	"github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/usecases"
	"github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/infrastructure/dependencies"
)

type ArticleHTTPHandler interface {
	// GetAllArticles method to return all articles.
	GetAllArticles(w http.ResponseWriter, r *http.Request)
	// PostArticles method to insert articles into DB.
	PostArticles(w http.ResponseWriter, r *http.Request)
}

type articleHTTPHandler struct {
	usecase usecases.ArticleUsecase
}

func NewArticleHandlerHTTP(container *dependencies.Container) ArticleHTTPHandler {
	usecase := usecases.NewArticleUsecase(container)
	return &articleHTTPHandler{
		usecase: usecase,
	}
}
