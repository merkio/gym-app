package tasks

import (
	loggerWrap "gym-app/common/logger"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = loggerWrap.NewLogger()
}
