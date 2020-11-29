package loghelpers

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// declare the object
var Logger *logrus.Logger

func init() {
	// initialization of logger
	Logger = logrus.New()

	// open a file
	f, err := os.OpenFile("logs/consumer.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	// Log as JSON instead of the default ASCII formatter.
	Logger.Formatter = &log.JSONFormatter{}

	// Output to stderr instead of stdout
	Logger.Out = f

	// warning severity
	Logger.Level = log.DebugLevel

}

//Info to print info
func Info(args ...interface{}) {
	Logger.Info(args)
}

//Error to print Errors
func Error(args ...interface{}) {
	Logger.Error(args)
}

//Warn to print Warnings
func Warn(args ...interface{}) {
	Logger.Warn(args)
}

//Debug to print Debug info
func Debug(args ...interface{}) {
	Logger.Debug(args)
}
