package log

import "github.com/sirupsen/logrus"

type logger struct {
	*logrus.Logger
}

type entry struct {
	*logrus.Entry
}
