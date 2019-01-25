package domain

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"

	"bytes"
	"github.com/spf13/viper"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"strings"
	"wood-serve/database"
	"wood-serve/entitys"
)

var (
	ErrBucket       = errors.New("Invalid bucket!")
	ErrInvalidImage = errors.New("Invalid image!")
)

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

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

func AddImage(filename string, creatorId string) error {
	db := database.GetDatabaseInstance()

	image := new(entitys.Image)
	image.CreatorId = creatorId
	image.Filename = filename

	return db.Connection.Create(image).Error
}

func SaveImage(image string, creatorId string) string {
	imagedir := viper.GetString("IMAGE_DIR")

	filenameHash := getMD5Hash(image)

	filename, err := saveImageToDisk(imagedir+filenameHash, image)

	AddImage(filename, creatorId)

	if err != nil {
		panic(err)
	}

	return filename
}
