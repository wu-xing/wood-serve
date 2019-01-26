package domain

import (
	"wood-serve/database"
	"wood-serve/entitys"
)

func AddArticleBox(userId string, name string) error {
	db := database.GetDatabaseInstance()

	box := new(entitys.ArticleBox)
	box.CreatorId = userId
	box.Name = name

	return db.Connection.Create(box).Error
}

func GetArticleBoxes(userId string) []entitys.ArticleBox {
	db := database.GetDatabaseInstance()

	articleBoxes := []entitys.ArticleBox{}

	db.Connection.Model(entitys.ArticleBox{}).Where("creator_id = ?", userId).Find(&articleBoxes)
	return articleBoxes
}
