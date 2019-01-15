package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	)

var(
	Log *logrus.Entry
)

func init(){
	logger := logrus.New()
	logger.Formatter = new(prefixed.TextFormatter)
	Log = logger.WithFields(logrus.Fields{"Mosaic": "Process"})
	Log.Logger.SetLevel(logrus.InfoLevel)
}
