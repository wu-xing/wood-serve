package main

import (
	"github.com/wu-xing/wood-serve/app/route"
	"github.com/wu-xing/wood-serve/database"
	"github.com/wu-xing/wood-serve/database/migration"
	"github.com/wu-xing/wood-serve/domain/config"
	"github.com/wu-xing/wood-serve/entitys"
	"github.com/wu-xing/wood-serve/log"
	"github.com/wu-xing/wood-serve/schedulers"
	"github.com/wu-xing/wood-serve/util"

	"github.com/robfig/cron"

	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/labstack/echo"
)

func setupReadConfg() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.SetEnvPrefix("WOOD")
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {
	util.LogBanner()

	setupReadConfg()

	log.SetupLogger()

	db := database.GetDatabaseInstance()
	db.Connection.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	defer db.Connection.Close()

	db.Connection.AutoMigrate(&entitys.ArticleBox{}, &entitys.Article{}, &entitys.User{}, &entitys.ArticleHistory{}, &entitys.Image{}, &config.AppConfig{})
	migration.AppMigration()

	e := echo.New()

	route.InitRoutes(e)

	c := cron.New()

	c.AddFunc("0 50 0 * * *", func() { // every day 1 am 50
		t := time.Now()
		fmt.Println("开始执行历史文章归档定时任务")
		fmt.Println(t.Format("2006-01-02 15:04:05"))
		schedulers.LogArticleHistory()
	})

	c.Start()

	fmt.Println("jellyfish serve on http://0.0.0.0:8020")

	t := time.Now()
	fmt.Println("🔥  WOOD SERVER LUNCHED  🔥")
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	e.Logger.Fatal(e.Start("0.0.0.0:" + "8020"))
}
