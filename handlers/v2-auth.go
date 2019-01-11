package handlers

import (
	"database/sql"
	"time"

	"github.com/dgrijalva/jwt-go"
	"net/http"
	"wood-serve/models"

	"github.com/dchest/captcha"
	"github.com/spf13/viper"

	// "fmt"

	"github.com/labstack/echo"
)

func V2SignUp(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{}

		request := new(struct {
			Captcha   string `json:"captcha"`
			CaptchaId string `json:"captchaId"`
			Username  string `json:"username"`
			Password  string `json:"password"`
		})

		c.Bind(&request)
		user.Username = request.Username
		user.Password = request.Password

		disableSignUp := viper.GetBool("DISABLE_SIGN_UP")

		if disableSignUp {
			return c.NoContent(http.StatusForbidden)
		}

		if !captcha.VerifyString(request.CaptchaId, request.Captcha) {
			return c.NoContent(http.StatusBadRequest)
		} else {
			_, error := models.CreateUser(db, &user)
			if error == nil {
				return c.NoContent(http.StatusNoContent)
			} else {
				return error
			}
		}

	}
}

func V2SignInV2(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		request := new(struct {
			Username string `json:"username"`
			Password string `json:"password"`
		})

		c.Bind(&request)

		isExist := models.CheckUserExist(db, request.Username)

		if !isExist {
			return c.JSON(http.StatusBadRequest, "User do not exist")
		}

		user, err := models.GetUserWhenCompareHashAndPassword(db, request.Username, request.Password)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "")
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Username
		claims["id"] = user.ID
		claims["createdAt"] = user.CreatedAt
		claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

		// TODO replace secret
		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
			"id":    user.ID,
		})
	}
}
