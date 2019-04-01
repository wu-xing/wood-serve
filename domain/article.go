package domain

import (
	"github.com/fwchen/wood-serve/database"
	"github.com/fwchen/wood-serve/entitys"
	log "github.com/fwchen/wood-serve/log"
	"time"
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

func SearchArticleByUserId(userId string, matchStr string) []entitys.Article {
	db := database.GetDatabaseInstance()

	articles := []entitys.Article{}

	db.Connection.Where("creator_id = ? AND content like ?", userId, "%"+matchStr+"%").Find(&articles)
	return articles
}

func CheckArticleBelongUser(userId string, articleId string) bool {
	db := database.GetDatabaseInstance()
	article := new(entitys.Article)
	db.Connection.Where("creator_id = ? AND id = ?", userId, articleId).First(&article)
	return article.ID != ""
}

func GetArticlesByUser(userId string) []entitys.Article {
	db := database.GetDatabaseInstance()
	articles := []entitys.Article{}
	db.Connection.Where("creator_id = ?", userId).Find(&articles)

	log.EventLog.Info("user id [" + userId + "] get articles")

	return articles
}

func GetArticleByID(articleId string) (entitys.Article, error) {
	db := database.GetDatabaseInstance()

	article := entitys.Article{}
	error := db.Connection.Where("id = ?", articleId).First(&article).Error
	return article, error
}

func GetArticleByIdAndUserId(userId string, articleId string) (entitys.Article, error) {
	db := database.GetDatabaseInstance()

	article := entitys.Article{}
	error := db.Connection.Where("id = ? AND creator_id = ?", articleId, userId).First(&article).Error
	return article, error
}

func GetAllArticleModifyTody() ([]entitys.Article, error) {
	db := database.GetDatabaseInstance()

	t := time.Now()
	subADay, _ := time.ParseDuration("-24h")
	addADay, _ := time.ParseDuration("+24h")
	yestoday := t.Add(subADay)
	local, _ := time.LoadLocation("Asia/Shanghai")
	yestodayZeroTime := yestoday.Format("2006-01-02")

	compareBegin, _ := time.ParseInLocation("2006-01-02", yestodayZeroTime, local)
	compareEnd := compareBegin.Add(addADay)

	articles := []entitys.Article{}
	error := db.Connection.Where("updated_at < ? ::date AND updated_at > ? ::date", compareEnd, compareBegin).Find(&articles).Error
	return articles, error
}

func LetArticleShare(userId string, articleId string) error {
	db := database.GetDatabaseInstance()

	article := new(entitys.Article)

	return db.Connection.Model(article).Where("creator_id = ? And id = ?", userId, articleId).Update("is_public", true).Error
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
