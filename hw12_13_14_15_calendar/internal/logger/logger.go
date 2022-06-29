package logger

import (
	"github.com/sirupsen/logrus"
)

type WrapLogger struct {
	log *logrus.Logger
}

func New(level string) *WrapLogger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return &WrapLogger{
		log: logger,
	}
}

func (l *WrapLogger) Info(msg string) {
	l.log.Info(msg)
}

func (l *WrapLogger) Error(msg string) {
	l.log.Error(msg)
}
