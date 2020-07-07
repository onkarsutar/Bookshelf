package logging

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)


//Logger - Global instance of logger
var logger *logrus.Logger

// InitLogging init logging
func InitLogging(){

    logger = logrus.New()
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal("Log OpenFile: ",err)
    }

    defer logFile.Close()

    
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
    logger.Out = os.Stdout

}

//LogError - Method to log errors
func LogError(args ...interface{}) {
	logger.Error(args)
}

//LogInfo - Method to log errors
func LogInfo(args ...interface{}) {
	logger.Info(args)
}

//LogWran - Method to log errors
func LogWran(args ...interface{}) {
	logger.Warn(args)
}

//LogDebug - Method to log errors
func LogDebug(args ...interface{}) {
	logger.Debug(args)
}