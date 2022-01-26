package helper

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()

	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	PanicIfError(err)

	logger.SetOutput(file)
}

func LoggerInit() *logrus.Logger {
	return logger
}
