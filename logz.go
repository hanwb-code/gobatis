package gobatis

import (
	"go.uber.org/zap"
)

func init() {
}

func NewStdLogger() *LoggerImpl {
	logger, _ := zap.NewProduction()

	return &LoggerImpl{
		logger: logger.Sugar(),
	}
}

type LoggerImpl struct {
	logger *zap.SugaredLogger
}

func (c *LoggerImpl) SetLevel(level Level) {

}
func (c *LoggerImpl) Sync() error {

	return nil
}
func (c *LoggerImpl) Fatalf(format string, args ...interface{}) {
	c.logger.Fatalf(format, args...)
}
func (c *LoggerImpl) Errorf(format string, args ...interface{}) {
	c.logger.Errorf(format, args...)
}
func (c *LoggerImpl) Panicf(format string, args ...interface{}) {
	c.logger.Panicf(format, args...)
}
func (c *LoggerImpl) Warnf(format string, args ...interface{}) {
	c.logger.Warnf(format, args...)
}
func (c *LoggerImpl) Infof(format string, args ...interface{}) {
	c.logger.Infof(format, args...)
}
func (c *LoggerImpl) Debugf(format string, args ...interface{}) {
	c.logger.Debugf(format, args...)
}
