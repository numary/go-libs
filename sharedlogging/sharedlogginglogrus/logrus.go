package sharedlogginglogrus

import (
	"github.com/numary/go-libs/sharedlogging"
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	entry interface {
		Debugf(format string, args ...interface{})
		Debug(args ...interface{})
		Infof(format string, args ...interface{})
		Info(args ...interface{})
		Errorf(format string, args ...interface{})
		Error(args ...interface{})
		WithFields(fields logrus.Fields) *logrus.Entry
	}
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}
func (l *logrusLogger) Debugf(fmt string, args ...interface{}) {
	l.entry.Debugf(fmt, args...)
}
func (l *logrusLogger) Infof(fmt string, args ...interface{}) {
	l.entry.Infof(fmt, args...)
}
func (l *logrusLogger) Info(args ...interface{}) {
	l.entry.Info(args...)
}
func (l *logrusLogger) Errorf(fmt string, args ...interface{}) {
	l.entry.Errorf(fmt, args...)
}
func (l *logrusLogger) Error(args ...interface{}) {
	l.entry.Error(args...)
}
func (l *logrusLogger) WithFields(fields map[string]interface{}) sharedlogging.Logger {
	return &logrusLogger{
		entry: l.entry.WithFields(fields),
	}
}

var _ sharedlogging.Logger = &logrusLogger{}

func New(logger *logrus.Logger) *logrusLogger {
	return &logrusLogger{
		entry: logger,
	}
}