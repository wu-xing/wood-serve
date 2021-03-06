package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	log "github.com/wu-xing/wood-serve/log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	Connection *gorm.DB
}

var dbInstance *DB

func GetDatabaseInstance() *DB {
	if dbInstance == nil {
		dbInstance = new(DB)
		dbInstance.Connection = ConnectDatabase()
	}
	return dbInstance
}

func ConnectDatabase() *gorm.DB {
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbName := viper.GetString("DB_NAME")
	dbPasswd := viper.GetString("DB_PASSWD")

	connectConfig := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPasswd)

	db, err := gorm.Open("postgres", connectConfig)

	db.SetLogger(log.DBLog)
	db.LogMode(true)

	if err != nil {
		panic(err)
	}
	return db
}
