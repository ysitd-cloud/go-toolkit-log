package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

func (l *entry) WithField(key string, value interface{}) Logger {
	e := l.Entry.WithField(key, value)
	return &entry{
		Entry: e,
	}
}

func (l *entry) SetFormatter(formatter logrus.Formatter) {
	l.Logger.Formatter = formatter
}

func (l *entry) SetLevel(level Level) {
	l.Level = level
}

func (l *entry) SetOutput(writer io.Writer) {
	l.Logger.Out = writer
}

