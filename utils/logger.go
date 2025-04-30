package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log is the global logger instance
var Log = logrus.New()

func InitLogger() {
	// Set log format as JSON
	Log.SetFormatter(&logrus.JSONFormatter{})

	// Set output to stdout (console)
	Log.SetOutput(os.Stdout)

	// Set log level (change to logrus.DebugLevel for more detailed logs)
	Log.SetLevel(logrus.InfoLevel)
}
