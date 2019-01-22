package domain

import (
	"wood-serve/database"
	"wood-serve/entitys"
)

func CreateArticle(title string, content string, userId string) (*entitys.Article, error) {
	db := database.GetDatabaseInstance()

	article := new(entitys.Article)
	article.CreatorId = userId
	article.Content = content
	article.Title = title

	error := db.Connection.Create(&article).Error
	return article, error
}
