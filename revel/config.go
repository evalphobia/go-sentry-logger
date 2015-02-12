package log_revel

import (
	logger "github.com/evalphobia/go-sentry-logger"
	conf "github.com/evalphobia/revel-config-loader"
)

const logConfigFileName = "log"

// override loggers in initialize
func init() {
	logger.SetConfig(NewRevelConfig())
}

type RevelConfig struct{}

func NewRevelConfig() *RevelConfig {
	return &RevelConfig{}
}

// retrieve value from revel's config file
func (c *RevelConfig) GetConfigValueDefault(section, key, df string) string {
	return conf.GetConfigValueDefault(logConfigFileName, section, section+"."+key, df)
}

func (c *RevelConfig) GetConfigValue(section, key string) string {
	return c.GetConfigValueDefault(section, key, "")
}
