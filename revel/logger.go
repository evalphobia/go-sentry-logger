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
	setRevelLogger(LOG.Logger)
}

// set loggers to revel's logger
func setRevelLogger(loggers map[string]*log.Logger) {
	loggers["fatal"] = revel.ERROR
	loggers["error"] = revel.ERROR
	loggers["warn"] = revel.WARN
	loggers["info"] = revel.INFO
	loggers["debug"] = revel.TRACE
}
