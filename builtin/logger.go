package builtin

import "github.com/sirupsen/logrus"

var LOGGER *logrus.Logger

func init() {
	LOGGER = logrus.New()
	LOGGER.SetFormatter(&logrus.TextFormatter{})
}
