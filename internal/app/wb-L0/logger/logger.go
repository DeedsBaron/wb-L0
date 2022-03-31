package logger

import (
	"github.com/sirupsen/logrus"
	"log"
)

var (
	Log *logrus.Logger
)

func ConfigureLogger(logLevel string) {
	Log = logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatalln("Cant configure logger: ", err.Error())
	}
	Log.SetLevel(level)
}
