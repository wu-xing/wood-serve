package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing Task data
type Article struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	IsPublic     bool
	IsEncryption bool   `json:"isEncryption"`
	Status       string `json:"status"`
	CreaterId    string `json:"createrId"`
	CreatedAt    int64  `json:"createdAt"`
	UpdateAt     int64  `json:"updatedAt"`
}

type ArticleHistory struct {
	ArticleId string `json:"articleId"`
	Content   string `json:"content"`
	Date      string `json:"date"`
}

type ArticleCollection struct {
	Articles []Article `json:"items"`
}

func LetArticleEncryption(db *sql.DB, articleId string) (int64, error) {
	sql := "UPDATE articles set is_encryption = 1 where id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(articleId)

	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

func SearchArticle(db *sql.DB, userId string, searchStr string) ArticleCollection {
	sqlstr := "SELECT id, content, title, status, is_encryption, created_at, updated_at FROM articles where creater_id = ? and content like '%' || ? || '%'"
	rows, err := db.Query(sqlstr, userId, searchStr)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	articleCollection := ArticleCollection{Articles: make([]Article, 0)}

	for rows.Next() {
		article := Article{}
		// TODO 判断空条件
		var createdAt time.Time
		var updatedAt time.Time
		var isEncryption sql.NullBool
		err2 := rows.Scan(&article.ID, &article.Content, &article.Title, &article.Status, &isEncryption, &createdAt, &updatedAt)
		article.CreatedAt = createdAt.UnixNano() / int64(time.Millisecond)
		article.UpdateAt = updatedAt.UnixNano() / int64(time.Millisecond)
		article.IsEncryption = isEncryption.Bool

		if err2 != nil {
			panic(err2)
		}
		articleCollection.Articles = append(articleCollection.Articles, article)
	}
	return articleCollection
}

func LetArticleShare(db *sql.DB, articleId string) (int64, error) {
	sql := "UPDATE articles set is_public = 1 where id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(articleId)

	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

func GetArticlesFromDB(db *sql.DB, userId string) ArticleCollection {
	sqlstr := "SELECT id, content, title, status, is_encryption, created_at, updated_at FROM articles where creater_id = ?"
	rows, err := db.Query(sqlstr, userId)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	articleCollection := ArticleCollection{Articles: make([]Article, 0)}

	for rows.Next() {
		article := Article{}
		// TODO 判断空条件
		var createdAt time.Time
		var updatedAt time.Time
		var isEncryption sql.NullBool
		err2 := rows.Scan(&article.ID, &article.Content, &article.Title, &article.Status, &isEncryption, &createdAt, &updatedAt)
		article.CreatedAt = createdAt.UnixNano() / int64(time.Millisecond)
		article.UpdateAt = updatedAt.UnixNano() / int64(time.Millisecond)
		article.IsEncryption = isEncryption.Bool

		if err2 != nil {
			panic(err2)
		}
		articleCollection.Articles = append(articleCollection.Articles, article)
	}
	return articleCollection
}

func CheckArticleBelong(db *sql.DB, articleId string, userId string) bool {
	sql := "SELECT id FROM articles WHERE id = ? and creater_id = ?"
	row := db.QueryRow(sql, articleId, userId)

	var id string
	err := row.Scan(&id)
	return err == nil
}

func GetArticle(db *sql.DB, articleId string) Article {
	sqlstr := "SELECT id, content, status, creater_id, created_at, updated_at, is_public, is_encryption FROM articles where id = ?"
	row := db.QueryRow(sqlstr, articleId)

	var article Article
	var createdAt time.Time
	var updatedAt time.Time
	var isPublic sql.NullBool
	var isEncryption sql.NullBool

	err := row.Scan(&article.ID, &article.Content, &article.Status, &article.CreaterId, &createdAt, &updatedAt, &isPublic, &isEncryption)
	article.CreatedAt = createdAt.UnixNano() / int64(time.Millisecond)
	article.UpdateAt = updatedAt.UnixNano() / int64(time.Millisecond)
	article.IsPublic = isPublic.Bool
	article.IsEncryption = isEncryption.Bool

	if err != nil {
		panic(err)
	}
	return article
}

func GetArticleHistoryByDate(db *sql.DB, articleId string, date string) ArticleHistory {
	sql := "select articles_history.article_id, articles_history.content, articles_history.date from articles, articles_history where articles.id = articles_history.article_id and articles.id = ? and articles_history.date = ?"
	row := db.QueryRow(sql, articleId, date)

	var articleHistory ArticleHistory
	err := row.Scan(&articleHistory.ArticleId, &articleHistory.Content, &articleHistory.Date)
	if err != nil {
		panic(err)
	}
	return articleHistory
}

func GetArticleHistoryDays(db *sql.DB, articleId string) []string {
	sql := "SELECT date FROM articles_history WHERE article_id = ?"
	rows, err := db.Query(sql, articleId)
	defer rows.Close()

	dayCollection := make([]string, 0)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var day string
		err2 := rows.Scan(&day)

		if err2 != nil {
			panic(err2)
		}
		dayCollection = append(dayCollection, day)
	}
	return dayCollection
}

func UpdateArticle(db *sql.DB, article *Article) (int64, error) {
	sql := "UPDATE articles set content = ?, title = ?, updated_at = ? where id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(article.Content, article.Title, time.Now().UnixNano()/int64(time.Millisecond), article.ID)

	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

// PutTask into DB
func CreateArticle(db *sql.DB, article *Article) (int64, error) {
	sql := "INSERT INTO articles(content, title, creater_id, status, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	result, err2 := stmt.Exec(article.Content, article.Title, article.CreaterId, article.Status, time.Now().UnixNano()/int64(time.Millisecond), time.Now().UnixNano()/int64(time.Millisecond))

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// DeleteTask from DB
func DeleteArticle(db *sql.DB, id int, userId string) (int64, error) {
	sql := "DELETE FROM articles WHERE id = ? and creater_id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id, userId)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
