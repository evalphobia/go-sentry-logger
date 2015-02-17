/*
	for revel setting
*/

package log_revel

import (
	LOG "github.com/evalphobia/go-sentry-logger"
	"github.com/revel/revel"
	"log"
)

// override loggers in initialize
func init() {
	LOG.Logger = &RevelLogger{
		loggers: make(map[string]*log.Logger),
	}
}

type RevelLogger struct {
	loggers map[string]*log.Logger
}

func (l *RevelLogger) GetLogger(name string) *log.Logger {
	logger, ok := l.loggers[name]
	if !ok {
		switch name {
		case "fatal":
			logger = revel.ERROR
		case "error":
			logger = revel.ERROR
		case "warn":
			logger = revel.WARN
		case "info":
			logger = revel.INFO
		case "debug":
			logger = revel.TRACE
		default:
			logger = revel.TRACE
		}
		l.loggers[name] = logger
	}
	return logger
}
