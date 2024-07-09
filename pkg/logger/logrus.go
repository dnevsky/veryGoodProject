package logger

import (
	"github.com/sirupsen/logrus"
)

type LogManager struct {
}

func NewLogManager() *LogManager {
	// null, _ := os.OpenFile(os.DevNull, os.O_RDWR, 0666)
	// logrus.SetOutput(null)

	return &LogManager{}
}

func (m *LogManager) Debug(msg string, params ...interface{}) {
	logrus.Debug(msg, params)
}

func (m *LogManager) Info(msg string, params ...interface{}) {
	logrus.Info(msg, params)
}

func (m *LogManager) Warn(msg string, params ...interface{}) {
	logrus.Warn(msg)
}

func (m *LogManager) Error(err error, params ...interface{}) {
	logrus.WithError(err).Error(err, params)
}
