package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	// Config fields e.g. port or username
	Port string `yaml:"port"`
}

func LoadConfig(path string) (config Config, err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	err = yaml.Unmarshal(b, &config)
	return
}
