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

func GetArticlesByUser(userId string) []entitys.Article {

	db := database.GetDatabaseInstance()

	articles := []entitys.Article{}
	db.Connection.Where("creator_id = ?", userId).Find(&articles)
	return articles
}
