package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"shopping-bot/config"
)

type Config struct {
	Bot struct {
		Token string `yaml:"token"`
	} `yaml:"bot"`
}

func NewConfig() *Config {
	c := &Config{}

	cFile := config.GetConfigFile()
	err := yaml.Unmarshal(cFile, c)

	if err != nil {
		log.Fatalln(err)
	}

	return c
}
