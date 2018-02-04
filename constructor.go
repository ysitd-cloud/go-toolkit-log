package log

import "github.com/sirupsen/logrus"

func New() Logger {
	return &logger {
		Logger: logrus.New(),
	}
}