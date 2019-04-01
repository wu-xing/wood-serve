package log

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// DBLog : db logger
var DBLog = logrus.New()

func setupDbLogger() {
	dbLogFile := viper.GetString("DB_LOG_FILE")

	DBLog.SetOutput(os.Stdout)

	file, err := os.OpenFile(dbLogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		DBLog.SetOutput(file)
	} else {
		DBLog.Info("Failed to log to file, using default stderr")
	}
}
