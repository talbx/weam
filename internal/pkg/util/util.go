package util

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	ApiKey   string `yaml:"apiKey"`
	Receiver string `yaml:"receiver"`
}

var AppConf AppConfig

func ReadConfig(config *AppConfig) {
	filename := "config.yaml"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(buf, config)
	AppConf = *config
	if err != nil {
		panic(err)
	}
}

func MapToString(m map[string]string) string {
	var tplString = ""
	for key, value := range m {
		tplString += fmt.Sprintf(`- %s: %s \n`, key, value)
	}
	return tplString
}
