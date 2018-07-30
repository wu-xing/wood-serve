package main

import (
	"database/sql"

	"net/http"
	"wood-serve/database"
	"wood-serve/handlers"
	"wood-serve/schedulers"

	"github.com/robfig/cron"

	"fmt"
	"github.com/dchest/captcha"
	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
	"time"

	"bytes"
	_ "github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	// "os"

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
        is_encryption BOOL,
        created_at DATE,
        updated_at DATE
    );

    CREATE TABLE IF NOT EXISTS articles_history (
        article_id INTEGER NOT NULL,
        date CHAR(20) NOT NULL,
        content TEXT NOT NULL,
        PRIMARY KEY(article_id, date)
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
	isEnabled := true
	isColorEnabled := true
	banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, bytes.NewBufferString(" ▄▀▀▄    ▄▀▀▄  ▄▀▀▀▀▄   ▄▀▀▀▀▄   ▄▀▀█▄▄\n█   █    ▐  █ █      █ █      █ █ ▄▀   █ \n▐  █        █ █      █ █      █ ▐ █    █\n█   ▄    █  ▀▄    ▄▀ ▀▄    ▄▀   █    █\n▀▄▀ ▀▄ ▄▀    ▀▀▀▀     ▀▀▀▀    ▄▀▄▄▄▄▀ \n▀                      █     ▐  \n▐"))
	fmt.Println("")

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
	r.POST("/article/encryption/:articleId", handlers.LetArticleEncryption(db))

	r.GET("/article/:articleId/history", handlers.GetArticleHistory(db))

	r.POST("/image/base64", handlers.PostAvatarByBase64(db))

	fmt.Println("jellyfish serve on http://0.0.0.0:8020")

	c := cron.New()
	c.AddFunc("0 50 0 * * *", func() { // every day 1 am
		t := time.Now()
		fmt.Println("开始执行历史文章归档定时任务")
		fmt.Println(t.Format("2006-01-02 15:04:05"))
		schedulers.LogArticleHistory(db)
	})
	c.Start()

	t := time.Now()
	fmt.Println("Wood 服务启动")
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	e.Logger.Fatal(e.Start("0.0.0.0:8020"))
}
