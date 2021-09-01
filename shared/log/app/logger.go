package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func Init(fns ...func(l *logrus.Logger)) {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	for _, fn := range fns {
		fn(logger)
	}
}

func Trace(fields *logrus.Fields, args ...interface{}) {
	writeToLog(logrus.TraceLevel, fields, args...)
}

func Debug(fields *logrus.Fields, args ...interface{}) {
	writeToLog(logrus.DebugLevel, fields, args...)
}

func Info(fields *logrus.Fields, args ...interface{}) {
	writeToLog(logrus.InfoLevel, fields, args...)
}

func Warn(fields *logrus.Fields, args ...interface{}) {
	writeToLog(logrus.WarnLevel, fields, args...)
}

func Error(fields *logrus.Fields, args ...interface{}) {
	writeToLog(logrus.ErrorLevel, fields, args...)
}

func Fatal(fields *logrus.Fields, args ...interface{}) {
	writeToLog(logrus.FatalLevel, fields, args...)
}

func Panic(fields *logrus.Fields, args ...interface{}) {
	writeToLog(logrus.PanicLevel, fields, args...)
}

func writeToLog(level logrus.Level, fields *logrus.Fields, args ...interface{}) {
	if fields != nil {
		logger.WithFields(*fields).Log(level, args...)
	} else {
		logger.Log(level, args...)
	}
}
