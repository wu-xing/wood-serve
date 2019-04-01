package database

import (
	"database/sql"
	"fmt"
	log "github.com/fwchen/wood-serve/log"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

type DatabaseInstance struct {
	Connection *gorm.DB
}

var dbInstance *DatabaseInstance

func GetDatabaseInstance() *DatabaseInstance {
	if dbInstance == nil {
		dbInstance = new(DatabaseInstance)
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
