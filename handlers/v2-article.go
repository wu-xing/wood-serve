package handlers

import (
	"database/sql"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"wood-serve/domain"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func V2GetArticles(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.QueryParam("userId")

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		jwtUserId := claims["id"].(string)

		if userId != jwtUserId {
			return c.NoContent(http.StatusUnauthorized)
		}
		articles := domain.GetArticlesByUser(userId)
		return c.JSON(http.StatusOK, articles)
	}
}

func V2GetShareArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		articleId := c.Param("id")

		article, error := domain.GetArticleById(articleId)

		if error != nil {
			return error
		}

		if !article.IsPublic || article.IsEncryption {
			return c.NoContent(http.StatusUnauthorized)
		}

		return c.JSON(http.StatusOK, article)
	}
}

func V2GetHistoryArticleByDate() echo.HandlerFunc {
	return func(c echo.Context) error {
		date := c.Param("date")
		articleId := c.Param("articleId")

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		isBelong := domain.CheckArticleBelongUser(userId, articleId)

		if !isBelong {
			return c.NoContent(http.StatusUnauthorized)
		}

		article := domain.GetArticleHistory(articleId, date)
		return c.JSON(http.StatusOK, article)

	}
}

func V2GetArticleHistory() echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		articleId := c.Param("articleId")

		_, error := domain.GetArticleByIdAndUserId(userId, articleId)

		if error != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		days := domain.GetArticeHistoryDatesById(articleId)
		return c.JSON(http.StatusOK, days)
	}
}

func V2LetArticleEncryption() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		articleId := c.Param("articleId")

		error := domain.LetArticleEncryption(userId, articleId)

		if error == nil {
			return c.NoContent(http.StatusCreated)
		} else {
			return error
		}
	}
}

func V2SearchArticleByMatch() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		searchStr := c.Param("search")
		articles := domain.SearchArticleByUserId(userId, searchStr)
		return c.JSON(http.StatusOK, articles)
	}
}

func V2LetArticleShare() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		articleId := c.Param("articleId")

		err := domain.LetArticleShare(userId, articleId)
		if err == nil {
			return c.NoContent(http.StatusOK)
		} else {
			return err
		}
	}
}

func V2PostArticle(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		request := new(struct {
			Content string `json:"content"`
			Title   string `json:"title"`
		})

		c.Bind(&request)

		article, err := domain.CreateArticle(request.Title, request.Content, userId)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"id": article.ID,
			})
		} else {
			return err
		}
	}
}

func V2UpdateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		request := new(struct {
			ID      string `json:"id"`
			Title   string `json:"title"`
			Content string `json:"content"`
		})

		c.Bind(&request)

		error := domain.UpdateArticle(userId, request.ID, request.Title, request.Content)

		if error == nil {
			return c.NoContent(http.StatusCreated)
		} else {
			return error
		}

	}
}
