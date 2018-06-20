package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing Task data
type Article struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	CreaterId string `json:"createrId"`
	CreatedAt int64  `json:"createdAt"`
}

// TaskCollection is collection of Tasks
type ArticleCollection struct {
	Articles []Article `json:"items"`
}

func GetArticlesFromDB(db *sql.DB, userId string) ArticleCollection {
	sql := "SELECT id, content, title, status, created_at FROM articles where creater_id = ?"
	rows, err := db.Query(sql, userId)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	articleCollection := ArticleCollection{Articles: make([]Article, 0)}

	for rows.Next() {
		article := Article{}
		var createdAt time.Time
		err2 := rows.Scan(&article.ID, &article.Content, &article.Title, &article.Status, &createdAt)
		article.CreatedAt = createdAt.UnixNano() / int64(time.Millisecond)

		if err2 != nil {
			panic(err2)
		}
		articleCollection.Articles = append(articleCollection.Articles, article)
	}
	return articleCollection
}

func GetArticle(db *sql.DB, articleId string) Article {
	sql := "SELECT id, content, status, creater_id, created_at FROM articles where id = ?"
	row := db.QueryRow(sql, articleId)

	var article Article
	err := row.Scan(&article.ID, &article.Content, &article.Status, &article.CreaterId, &article.CreatedAt)
	if err != nil {
		panic(err)
	}
	return article
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
	sql := "INSERT INTO articles(content, title, creater_id, status, created_at) VALUES(?, ?, ?, ?, ?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	result, err2 := stmt.Exec(article.Content, article.Title, article.CreaterId, article.Status, time.Now().UnixNano()/int64(time.Millisecond))

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
