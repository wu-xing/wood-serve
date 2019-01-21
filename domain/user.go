package domain

import (
	"wood-serve/database"
	"wood-serve/entitys"
)

func CheckUserExist() {

}

func AuthUser(username string, password string) {
	// db := database.GetDatabaseInstance()
}

func CreateUser(username string, password string) {
	db := database.GetDatabaseInstance()
	user := new(entitys.User)
	user.Username = username
	user.Hash = password
	db.Connection.Create(&user)
}
