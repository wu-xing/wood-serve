package models

import (
	"database/sql"
	"github.com/satori/go.uuid"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing Task data
type ArticleBox struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"createdAt"`
	UpdateAt  int64  `json:"updatedAt"`
}

// PutTask into DB
func CreateArticleBox(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO article_box(id, name, created_at, updated_at) VALUES(?, ?, ?, ?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	uuid := uuid.Must(uuid.NewV4())
	result, err2 := stmt.Exec(uuid, name, time.Now().UnixNano()/int64(time.Millisecond), time.Now().UnixNano()/int64(time.Millisecond))

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}
