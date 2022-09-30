package util

import (
	"fmt"
	"os"
)

type AppConfig struct {
	ApiKey   string `yaml:"apiKey"`
	Receiver string `yaml:"receiver"`
}

var AppConf AppConfig

func ReadConfig(config *AppConfig) {
	config.ApiKey = os.Getenv("API_KEY")
	config.Receiver = os.Getenv("RECEIVER")
}

func MapToString(m map[string]string) string {
	var tplString = ""
	for key, value := range m {
		tplString += fmt.Sprintf(`- %s: %s \n`, key, value)
	}
	return tplString
}
