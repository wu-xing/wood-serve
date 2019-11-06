package handlers

import (
	"github.com/wu-xing/wood-serve/domain"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func PostImageByBase64() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		request := new(struct {
			Image string `json:"image"`
		})

		c.Bind(&request)

		filename := domain.SaveImage(request.Image, userId)

		return c.JSON(http.StatusOK, map[string]string{
			"image": filename,
		})

	}
}
