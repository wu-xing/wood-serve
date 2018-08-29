package schedulers

import (
	"database/sql"

	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Article struct {
	ID      string
	Content string
}

type ArticleCollection struct {
	Articles []Article
}

func LogArticleHistory(db *sql.DB) {
	t := time.Now()
	subDay, _ := time.ParseDuration("-24h")
	yestoday := t.Add(subDay)
	local, _ := time.LoadLocation("Asia/Shanghai")
	tString := yestoday.Format("2006-01-02")
	tZero, _ := time.ParseInLocation("2006-01-02", tString, local)
	tBegin := tZero.UnixNano() / int64(time.Millisecond)
	tEnd := tBegin + 1000*60*60*24

	fmt.Println("tEnd", tEnd)
	fmt.Println("tBegin", tBegin)

	sql := "SELECT id, content FROM articles where updated_at > ? and updated_at < ?"
	rows, err := db.Query(sql, tBegin, tEnd)

	if err != nil {
		panic(err)
	}

	articleCollection := ArticleCollection{Articles: make([]Article, 0)}
	for rows.Next() {
		article := Article{}
		err2 := rows.Scan(&article.ID, &article.Content)
		if err2 != nil {
			panic(err2)
		}

		articleCollection.Articles = append(articleCollection.Articles, article)

	}

	defer rows.Close()

	for _, element := range articleCollection.Articles {
		insertHistorySql := "INSERT INTO articles_history(article_id, content, date) VALUES(?, ?, ?)"
		stmt, err := db.Prepare(insertHistorySql)
		if err != nil {
			panic(err)
		}

		defer stmt.Close()

		_, err2 := stmt.Exec(element.ID, element.Content, tString)
		if err2 != nil {
			panic(err2)
		}
		fmt.Println("article write history successful. article.id = ", element.ID, " date = ", tString)
	}
}
