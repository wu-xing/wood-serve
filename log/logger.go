package log

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func SetupLogger() {
	// with Json Formatter

	LOG_FILE := viper.GetString("LOG_FILE")

	log.SetOutput(os.Stdout)

	file, err := os.OpenFile(LOG_FILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

}
