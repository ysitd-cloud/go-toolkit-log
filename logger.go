package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

func (l *logger) WithField(key string, value interface{}) Logger {
	e := l.Logger.WithField(key, value)
	return &entry{
		Entry: e,
	}
}

func (l *logger) SetOutput(writer io.Writer) {
	l.Out = writer
}

func (l *logger) SetFormatter(formatter logrus.Formatter) {
	l.Formatter = formatter
}
