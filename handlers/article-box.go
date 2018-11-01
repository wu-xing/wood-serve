package handlers

import (
	"database/sql"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"wood-serve/models"

	"github.com/labstack/echo"
)

func PostArticleBox(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		request := new(struct {
			Name string `json:"name"`
		})

		c.Bind(&request)

		id := models.CreateArticleBox(db, userId, request.Name)

		return c.JSON(http.StatusCreated, H{
			"id": id,
		})
	}
}

func GetArticleBoxs(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		articleBoxs := models.GetArticleBoxs(db, userId).ArticleBoxs

		return c.JSON(http.StatusOK, articleBoxs)
	}
}
