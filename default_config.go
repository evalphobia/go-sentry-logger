package log

import (
	"encoding/json"
	"io/ioutil"
)

type DefaultConfig struct {
	rootPath string
	config   map[string]interface{}
}

func NewDefaultConfig(path string) *DefaultConfig {
	c := &DefaultConfig{}
	c.rootPath = path
	return c
}

func SetDefaultConfig() {
	SetConfig(NewDefaultConfig(""))
}

// get parameter from json file
func (c *DefaultConfig) GetConfigValueDefault(section, key, defaultValue string) string {
	if c.config == nil {
		file, e := ioutil.ReadFile(c.rootPath + "/" + logConfigFileName + ".json")
		if e != nil {
			errLog := NewLogData(e, 2)
			errLog.Label = "DefaultConfig File Load error"
			ErrorWithoutTrack(errLog)
		}
		json.Unmarshal(file, &c.config)
	}
	value := ParseToString(c.config[key])
	return value
}

func (c *DefaultConfig) GetConfigValue(section, key string) string {
	return c.GetConfigValueDefault(section, key, "")
}
