package main

import (
	"fmt"
	"github.com/wu-xing/wood-serve/database"
	"github.com/wu-xing/wood-serve/domain"
	"github.com/spf13/viper"
	"os"
)

func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.SetEnvPrefix("WOOD")
	viper.AutomaticEnv()
	viper.ReadInConfig() // Find and read the config file

	var username string = os.Args[1]
	var password string = os.Args[2]

	goDB := database.ConnectDatabase()
	goDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	error := domain.CreateUser(username, password)

	if error != nil {
		fmt.Print("Create user successful:")
		fmt.Print("username: ", username, "\npassword: ", password, "\n")
	}
}
