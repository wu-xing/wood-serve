package domain

import (
	"fmt"
	"time"
	"wood-serve/database"
	"wood-serve/entitys"
)

func AddArticleHistory(articleId string, date string, content string) error {
	db := database.GetDatabaseInstance()

	articleHistory := new(entitys.ArticleHistory)
	articleHistory.ArticleId = articleId
	articleHistory.Date = date
	articleHistory.Content = content

	return db.Connection.Create(articleHistory).Error
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
