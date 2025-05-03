package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type logFormatter struct {
	logrus.TextFormatter
}

func (f *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Customize the log format
	return []byte(fmt.Sprintf("[%s] [%s] %s\n", entry.Time.Format(time.RFC3339), entry.Level, entry.Message)), nil
}

// Log is the global logger instance
var Log = logrus.New()

func InitLogger() {
	// Set log format as JSON
	Log.SetFormatter(&logFormatter{})

	// Set output to stdout (console)
	Log.SetOutput(os.Stdout)

	// Set log level (change to logrus.DebugLevel for more detailed logs)
	Log.SetLevel(logrus.InfoLevel)
}
