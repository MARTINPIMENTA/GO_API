package database

import (
	"fmt"

	articleEntities "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/article/entities"
)

// Database mock
var dummyDB articleEntities.Articles

// GetAllArticlesFromDB gets articles from DB.
func (articleDatabase *articleDatabase) GetAllArticlesFromDB() (articleEntities.Articles, error) {
	articlesResponse := dummyDB
	if len(articlesResponse) == 0 {
		return nil, fmt.Errorf("error getting articles from DB")
	}

	// DB connection.
	// db, err := sql.Open("mysql", "root:root(127.0.0.1:8080/MK1Con)")
	// if err != nil {
	// 	return nil, err
	// }
	// // defer close function for DB when finished.
	// defer db.Close()

	return articlesResponse, nil

}

// PostArticleIntoDB inserts articles into DB.
func (articleDatabase *articleDatabase) PostArticleIntoDB(article articleEntities.Article) error {
	initialLen := len(dummyDB)
	dummyDB = append(dummyDB, article)
	finalLen := len(dummyDB)
	// Compare size of dummy DB to check for insert.
	if initialLen >= finalLen {
		return fmt.Errorf("error in DB insert for article")
	}

	return nil
}
