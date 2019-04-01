package log

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// EventLog : event logger
var EventLog = logrus.New()

func setupEventLogger() {
	evnetLogFile := viper.GetString("EVENT_LOG_FILE")

	EventLog.SetOutput(os.Stdout)

	file, err := os.OpenFile(evnetLogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		EventLog.SetOutput(file)
	} else {
		EventLog.Info("Failed to log to file, using default stderr")
	}
}
