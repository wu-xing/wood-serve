package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	// "time"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"wood-serve/domain"
	"wood-serve/models"

	"github.com/labstack/echo"
)

func V2GetArticles(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.QueryParam("userId")

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		jwtUserId := claims["id"].(string)

		if userId != jwtUserId {
			return c.JSON(http.StatusUnauthorized, "")
		}
		articles := models.GetArticlesFromDB(db, userId).Articles
		// for i := 0; i < len(articles); i++ {
		// 	if articles[i].IsEncryption.Valid {
		// 		articles[i].Content = ""
		// 	}
		// }
		return c.JSON(http.StatusOK, articles)
	}
}

func V2GetShareArticle(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		articleId := c.Param("id")

		article := models.GetArticle(db, articleId)
		fmt.Println(article.IsPublic)

		if !article.IsPublic || article.IsEncryption {
			return c.NoContent(http.StatusUnauthorized)
		}

		return c.JSON(http.StatusOK, article)
	}
}

func V2GetHistoryArticleByDate(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		date := c.Param("date")
		articleId := c.Param("articleId")

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		jwtUserId := claims["id"].(string)

		isBelong := models.CheckArticleBelong(db, articleId, jwtUserId)
		if !isBelong {
			return c.NoContent(http.StatusUnauthorized)
		}

		article := models.GetArticleHistoryByDate(db, articleId, date)
		return c.JSON(http.StatusOK, article)

	}
}

func V2GetArticleHistory(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		jwtUserId := claims["id"].(string)

		articleId := c.Param("articleId")
		isBelong := models.CheckArticleBelong(db, articleId, jwtUserId)
		if !isBelong {
			return c.NoContent(http.StatusUnauthorized)
		}

		days := models.GetArticleHistoryDays(db, articleId)
		return c.JSON(http.StatusOK, days)
	}
}

func V2PutArticle(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		article := new(models.Article)
		c.Bind(&article)

		_, err := models.UpdateArticle(db, article)
		if err == nil {
			return c.NoContent(http.StatusCreated)
		} else {
			return err
		}
	}
}

func V2LetArticleEncryption(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		articleId := c.Param("articleId")
		_, err := models.LetArticleEncryption(db, articleId)
		if err == nil {
			return c.NoContent(http.StatusOK)
		} else {
			return err
		}
	}
}

func V2SearchArticleByMatch(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		searchStr := c.Param("search")
		articles := models.SearchArticle(db, userId, searchStr)
		// for i := 0; i < len(articles); i++ {
		// 	if articles[i].IsEncryption.Valid {
		// 		articles[i].Content = ""
		// 	}
		// }
		return c.JSON(http.StatusOK, articles)
	}
}

func V2LetArticleShare(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		articleId := c.Param("articleId")
		_, err := models.LetArticleShare(db, articleId)
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

func V2DeleteArticle(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a task
		_, err := models.DeleteArticle(db, id, userId)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{})
			// Handle errors
		} else {
			return err
		}
	}
}