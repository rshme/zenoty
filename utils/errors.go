package utils

import (
	"github.com/rafiseptian90/zenoty/pkg/logger"
)

func FailOnError(msg string, err error) {
	logger.Log.Errorf("%v : %v", msg, err)
}
