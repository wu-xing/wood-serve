package route

import (
	"net/http"

	"github.com/dchest/captcha"
	"github.com/labstack/echo/middleware"
	_ "github.com/labstack/gommon/log"
	"github.com/wu-xing/wood-serve/handlers"

	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/upload", "upload")

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello my firends")
	})
	e.POST("/v2/signin", handlers.V2SignIn())
	e.POST("/v2/signup", handlers.V2SignUp())

	e.GET("/captcha/*", echo.WrapHandler(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
	e.POST("/captcha", handlers.GenCaptcha())

	e.GET("/v2/share/article/:id", handlers.V2GetShareArticle())

	r := e.Group("/auth")

	// TODO rename secret
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "header:App-Authorization",
	}))

	r.GET("/v2/articles", handlers.V2GetArticles())
	r.POST("/v2/article", handlers.V2PostArticle())

	r.PUT("/v2/article/:id", handlers.V2UpdateArticle())
	r.POST("/v2/article/encryption/:articleId", handlers.V2LetArticleEncryption())

	r.POST("/v2/article/share/:articleId", handlers.V2LetArticleShare())

	r.GET("/articles/search/:search", handlers.V2SearchArticleByMatch())

	r.GET("/article/:articleId/history", handlers.V2GetArticleHistory())
	r.GET("/article/:articleId/history/:date", handlers.V2GetHistoryArticleByDate())

	r.POST("/article-box", handlers.PostArticleBox())
	r.GET("/article-boxes", handlers.GetArticleBoxes())

	r.POST("/image/base64", handlers.PostImageByBase64())
}
