package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	logFileName := fmt.Sprintf("logs/zenoty-%v.log", time.Now().Format("2006-01-02"))
	logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		Log.Error(err)
	}
	Log.SetOutput(logFile)

	// Set output to JSON format
	Log.SetFormatter(&logrus.JSONFormatter{})

	// Set the log level to debug
	Log.SetLevel(logrus.DebugLevel)
}
