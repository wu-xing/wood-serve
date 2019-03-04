package main

import (
	"net/http"
	"wood-serve/database"
	"wood-serve/entitys"
	"wood-serve/handlers"
	"wood-serve/schedulers"
	"wood-serve/util"

	"github.com/robfig/cron"

	"fmt"
	"github.com/dchest/captcha"
	"time"

	_ "github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

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
	util.LogBanner()

	setupReadConfg()

	db := database.InitDB("storage.sqlite3?parseTime=true&cache=shared&mode=rwc")
	defer db.Close()

	goDB := database.ConnectDatabase()
	goDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	defer goDB.Close()

	goDB.AutoMigrate(&entitys.ArticleBox{}, &entitys.Article{}, &entitys.User{}, &entitys.ArticleHistory{}, &entitys.Image{})

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/upload", "upload")

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello my firends")
	})
	e.POST("/v2/signin", handlers.V2SignIn())
	e.POST("/v2/signup", handlers.V2SignUp())

	e.GET("/captcha/*", echo.WrapHandler(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
	e.POST("/captcha", handlers.GenCaptcha(db))

	e.GET("/v2/share/article/:id", handlers.V2GetShareArticle())

	r := e.Group("/auth")

	// TODO rename secret
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "header:App-Authorization",
	}))

	r.GET("/v2/articles", handlers.V2GetArticles(db))
	r.POST("/v2/article", handlers.V2PostArticle(db))
	r.PUT("/v2/article/:id", handlers.V2UpdateArticle())
	r.POST("/v2/article/encryption/:articleId", handlers.V2LetArticleEncryption())

	r.POST("/v2/article/share/:articleId", handlers.V2LetArticleShare())

	r.GET("/articles/search/:search", handlers.V2SearchArticleByMatch())

	r.GET("/article/:articleId/history", handlers.V2GetArticleHistory())
	r.GET("/article/:articleId/history/:date", handlers.V2GetHistoryArticleByDate())

	r.POST("/article-box", handlers.PostArticleBox())
	r.GET("/article-boxes", handlers.GetArticleBoxes())

	r.POST("/image/base64", handlers.PostImageByBase64(db))

	fmt.Println("jellyfish serve on http://0.0.0.0:8020")

	c := cron.New()

	c.AddFunc("0 50 0 * * *", func() { // every day 1 am 50
		t := time.Now()
		fmt.Println("开始执行历史文章归档定时任务")
		fmt.Println(t.Format("2006-01-02 15:04:05"))
		schedulers.LogArticleHistory()
	})

	c.Start()

	t := time.Now()
	fmt.Println("🔥  WOOD SERVER LUNCHED🔥")
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	e.Logger.Fatal(e.Start("0.0.0.0:" + "8020"))
}
