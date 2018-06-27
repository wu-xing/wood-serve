package handlers

import (
	"database/sql"

	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"bytes"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

var (
	ErrBucket       = errors.New("Invalid bucket!")
	ErrInvalidImage = errors.New("Invalid image!")
)

func saveImageToDisk(fileNameBase, data string) (string, error) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", ErrInvalidImage
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return "", err
	}
	// _, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	_, _, err = image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	if err != nil {
		return "", err
	}

	fileName := fileNameBase
	ioutil.WriteFile(fileName, buff.Bytes(), 0644)

	return fileName, err
}

func PostAvatarByBase64(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		request := new(struct {
			Image string `json:"image"`
		})

		c.Bind(&request)

		imagedir := viper.GetString("imagedir")

		fileNameHash := GetMD5Hash(request.Image)

		fileName, err := saveImageToDisk(imagedir+fileNameHash, request.Image)
		if err != nil {
			panic(err)
		}

		sql := "INSERT INTO images(creator, filename, created_at) VALUES(?, ?, ?)"

		stmt, err2 := db.Prepare(sql)

		if err2 != nil {
			panic(err2)
		}

		defer stmt.Close()
		_, err3 := stmt.Exec(userId, fileNameHash, time.Now().UnixNano()/int64(time.Millisecond))
		if err3 != nil {
			panic(err3)
		}

		return c.JSON(http.StatusOK, map[string]string{
			"image": fileName,
		})

	}
}
