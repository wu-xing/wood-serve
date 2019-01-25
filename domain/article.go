package domain

import (
	"time"
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

func GetAllArticleModifyTody() ([]entitys.Article, error) {
	db := database.GetDatabaseInstance()

	t := time.Now()
	subDay, _ := time.ParseDuration("-24h")
	yestoday := t.Add(subDay)
	local, _ := time.LoadLocation("Asia/Shanghai")
	tString := yestoday.Format("2006-01-02")
	tZero, _ := time.ParseInLocation("2006-01-02", tString, local)
	tBegin := tZero.UnixNano() / int64(time.Millisecond)
	tEnd := tBegin + 1000*60*60*24

	articles := []entitys.Article{}
	error := db.Connection.Where("updated_at = < ? AND updated_at > ?", tEnd, tBegin).Find(&articles).Error
	return articles, error
}

func UpdateArticle(userId string, articleId string, title string, content string) error {
	db := database.GetDatabaseInstance()

	article := new(entitys.Article)

	db.Connection.Where("creator_id = ? And id = ?", userId, articleId).First(&article)

	article.Title = title
	article.Content = content

	return db.Connection.Save(&article).Error
}

func LetArticleEncryption(userId string, articleId string) error {
	db := database.GetDatabaseInstance()

	article := new(entitys.Article)

	return db.Connection.Model(article).Where("creator_id = ? And id = ?", userId, articleId).Update("is_encryption", true).Error
}
