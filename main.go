package main

import (
	"database/sql"

	"net/http"
	"wood-serve/database"
	"wood-serve/entitys"
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
        is_public BOOL,
        created_at DATE,
        updated_at DATE
    );

    CREATE TABLE IF NOT EXISTS article_box (
        id CHAR(36) PRIMARY KEY,
        creater_id INTEGER NOT NULL,
        name VARCHAR(200),
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

func setupReadConfg() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.SetEnvPrefix("WOOD")
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {
	banner.Init(colorable.NewColorableStdout(), true, true, bytes.NewBufferString("\n\n██╗    ██╗ ██████╗  ██████╗ ██████╗ \n██║    ██║██╔═══██╗██╔═══██╗██╔══██╗\n██║ █╗ ██║██║   ██║██║   ██║██║  ██║\n██║███╗██║██║   ██║██║   ██║██║  ██║\n╚███╔███╔╝╚██████╔╝╚██████╔╝██████╔╝\n ╚══╝╚══╝  ╚═════╝  ╚═════╝ ╚═════╝ \n\n\n"))
	fmt.Println("")

	setupReadConfg()

	db := database.InitDB("storage.sqlite3?parseTime=true&cache=shared&mode=rwc")
	defer db.Close()

	goDB := database.ConnectDatabase()
	goDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	defer goDB.Close()

	goDB.AutoMigrate(&entitys.ArticleBox{}, &entitys.Article{}, &entitys.User{})

	migrate(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/upload", "upload")

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello my firends")
	})
	e.POST("/signin", handlers.SignIn(db))
	e.POST("/v2/signin", handlers.V2SignIn())
	e.POST("/signup", handlers.SignUp(db))
	e.POST("/v2/signup", handlers.V2SignUp())

	e.GET("/captcha/*", echo.WrapHandler(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
	e.POST("/captcha", handlers.GenCaptcha(db))

	e.GET("/share/article/:id", handlers.GetShareArticle(db))

	r := e.Group("/auth")
	r.Use(middleware.JWT([]byte("secret")))

	r.GET("/articles", handlers.GetArticles(db))
	r.GET("/v2/articles", handlers.V2GetArticles(db))
	r.POST("/article", handlers.PostArticle(db))
	r.POST("/v2/article", handlers.V2PostArticle(db))
	r.DELETE("/article/:id", handlers.DeleteArticle(db))
	r.PUT("/article/:id", handlers.PutArticle(db))
	r.POST("/article/encryption/:articleId", handlers.LetArticleEncryption(db))
	r.POST("/article/share/:articleId", handlers.LetArticleShare(db))
	r.GET("/articles/search/:search", handlers.SearchArticleByMatch(db))

	r.GET("/article/:articleId/history", handlers.GetArticleHistory(db))
	r.GET("/article/:articleId/history/:date", handlers.GetHistoryArticleByDate(db))

	r.POST("/article-box", handlers.PostArticleBox(db))
	r.GET("/article-box", handlers.GetArticleBoxs(db))

	r.POST("/image/base64", handlers.PostAvatarByBase64(db))

	fmt.Println("jellyfish serve on http://0.0.0.0:8020")

	c := cron.New()
	c.AddFunc("0 50 0 * * *", func() { // every day 1 am 50
		t := time.Now()
		fmt.Println("开始执行历史文章归档定时任务")
		fmt.Println(t.Format("2006-01-02 15:04:05"))
		schedulers.LogArticleHistory(db)
	})

	c.Start()

	t := time.Now()
	fmt.Println("🔥  WOOD SERVER LUNCHED🔥")
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	e.Logger.Fatal(e.Start("0.0.0.0:8020"))
}
