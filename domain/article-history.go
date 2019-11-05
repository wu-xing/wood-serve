package domain

import (
	"fmt"
	"github.com/wu-xing/wood-serve/database"
	"github.com/wu-xing/wood-serve/entitys"
	"time"
)

func AddArticleHistory(articleId string, date string, content string) error {
	db := database.GetDatabaseInstance()

	articleHistory := new(entitys.ArticleHistory)
	articleHistory.ArticleId = articleId
	articleHistory.Date = date
	articleHistory.Content = content

	return db.Connection.Create(articleHistory).Error
}

func GetArticeHistoryDatesById(articeId string) []string {
	db := database.GetDatabaseInstance()

	articleHistorys := []entitys.ArticleHistory{}
	db.Connection.Where("article_id = ?", articeId).Find(&articleHistorys)

	dates := []string{}

	for _, history := range articleHistorys {
		dates = append(dates, history.Date)
	}
	return dates
}

func GetArticleHistory(articleId string, date string) entitys.ArticleHistory {
	db := database.GetDatabaseInstance()

	articleHistory := entitys.ArticleHistory{}
	db.Connection.Where("article_id = ? AND date = ?", articleId, date).Find(&articleHistory)
	return articleHistory
}

func ArchiveArticleHistory() {
	t := time.Now()
	subDay, _ := time.ParseDuration("-24h")
	yestoday := t.Add(subDay)
	tString := yestoday.Format("2006-01-02")

	articles, error := GetAllArticleModifyTody()
	if error != nil {
		panic(error)
	}

	for _, element := range articles {

		error2 := AddArticleHistory(element.ID, tString, element.Content)

		if error2 != nil {
			panic(error)
		}

		fmt.Println("article write history successful. article.id = ", element.ID, " date = ", tString)
	}
}
