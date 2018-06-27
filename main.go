package main

import (
	"database/sql"

	"net/http"
	"wood-serve/database"
	"wood-serve/handlers"

	"fmt"
	"github.com/dchest/captcha"

	_ "github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS articles(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        creater_id INTEGER NOT NULL,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        status TEXT,
        created_at DATE,
        updated_at DATE
    );

    CREATE TABLE IF NOT EXISTS users(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        hash TEXT NOT NULL,
        created_at DATE,
        updated_at DATE
    );

    CREATE TABLE IF NOT EXISTS images(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        creator INTEGER NOT NULL,
        filename TEXT NOT NULL,
        created_at DATE
    );


    `
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}

func readConfg() {
	viper.SetConfigName("config") // name of config file (without extension)

	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {
	readConfg()

	db := database.InitDB("storage.sqlite3?parseTime=true&cache=shared&mode=rwc")
	defer db.Close()

	migrate(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/upload", "upload")

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello my firend")
	})
	e.POST("/signin", handlers.SignIn(db))
	e.POST("/signup", handlers.SignUp(db))

	e.GET("/captcha/*", echo.WrapHandler(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
	e.POST("/captcha", handlers.GenCaptcha(db))

	r := e.Group("/auth")
	r.Use(middleware.JWT([]byte("secret")))

	r.GET("/articles", handlers.GetArticles(db))
	r.POST("/article", handlers.PostArticle(db))
	r.DELETE("/article/:id", handlers.DeleteArticle(db))
	r.PUT("/article/:id", handlers.PutArticle(db))

	r.POST("/image/base64", handlers.PostAvatarByBase64(db))

	fmt.Println("jellyfish serve on http://0.0.0.0:8020")
	e.Logger.Fatal(e.Start("0.0.0.0:8020"))

}
