package main

import "github.com/sirupsen/logrus"

var LOGGER_MODULE *logrus.Logger

func init() {
	LOGGER_MODULE = logrus.New()
	LOGGER_MODULE.SetFormatter(&logrus.TextFormatter{})
}
