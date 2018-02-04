package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Level = logrus.Level
type Formatter = logrus.Formatter

type Logger interface {
	SetOutput(io.Writer)
	SetLevel(level Level)
	SetFormatter(formatter logrus.Formatter)

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
}

type WithField interface {
	WithField(key string, value interface{}) Logger
}

type FieldLogger interface {
	WithField
	Logger
}