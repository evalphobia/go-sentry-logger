/*
	this package enable logging data with sending to sentry.
	see: https://github.com/evalphobia/go-sentry-logger
*/

package log

import (
	"log"
)

const (
	logConfigFileName = "log"

	sentryLevelDebug = (iota + 1)
	sentryLevelInfo
	sentryLevelWarn
	sentryLevelError
	sentryLevelFatal
)

var (
	Logger      map[string]*log.Logger
	config      Config
	sentryLevel int
)

func init() {
	Logger = make(map[string]*log.Logger)
	SetDefaultLogger(Logger)
	SetDefaultConfig()
	sentryLevel = sentryLevelWarn
}

// get parameter from config file
func getConfigValue(section, key string) string {
	return config.GetConfigValue(section, key)
}

// get parameter from config file
func SetConfig(conf Config) {
	config = conf
}

type Config interface {
	// params(filename, section, key, defaultValue)
	GetConfigValueDefault(string, string, string) string

	// params(filename, section, key)
	GetConfigValue(string, string) string
}
