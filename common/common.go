package common

import "github.com/wu-xing/wood-serve/database"

func GetDB() *database.DB {
	return database.GetDatabaseInstance()
}
