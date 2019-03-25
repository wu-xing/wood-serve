package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/fwchen/wood-serve/domain"

	"github.com/labstack/echo"
)

func PostArticleBox() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		request := new(struct {
			Name string `json:"name"`
		})

		c.Bind(&request)

		error := domain.AddArticleBox(userId, request.Name)

		if error != nil {
			return error
		}
		return c.NoContent(http.StatusCreated)
	}
}

func GetArticleBoxes() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		articleBoxs := domain.GetArticleBoxes(userId)

		return c.JSON(http.StatusOK, articleBoxs)
	}
}
