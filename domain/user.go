package domain

import (
	"wood-serve/database"
	"wood-serve/entitys"

	"golang.org/x/crypto/bcrypt"
)

func CheckUserExist(username string) bool {
	db := database.GetDatabaseInstance()

	user := new(entitys.User)
	return db.Connection.Where(&entitys.User{Username: username}).First(&user).Error == nil
}

func AuthUser(username string, password string) (*entitys.User, error) {
	db := database.GetDatabaseInstance()
	user := new(entitys.User)
	error := db.Connection.Where(&entitys.User{Username: username}).First(&user).Error
	if error != nil {
		return user, error

	}
	error2 := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
	return user, error2
}

func CreateUser(username string, password string) error {
	db := database.GetDatabaseInstance()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	user := new(entitys.User)
	user.Username = username
	user.Hash = string(hash)
	return db.Connection.Create(&user).Error
}
