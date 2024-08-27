package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	GetLogger() *logrus.Logger
	LogInfo(string)
	LogError(string, error)
}

type loggerImpl struct {
	logger *logrus.Logger
}

var globalLogger *logrus.Logger

func NewLogger() Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	globalLogger = logger
	return &loggerImpl{logger: logger}
}

func (l *loggerImpl) GetLogger() *logrus.Logger {
	return globalLogger
}

func (l *loggerImpl) LogInfo(message string) {
	l.logger.Info(message)
}

func (l *loggerImpl) LogError(message string, err error) {
	l.logger.WithError(err).Error(message)
}
