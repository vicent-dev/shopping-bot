package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const filename = "config.yaml"

type Config struct {
	Bot struct {
		Token string `yaml:"token"`
	} `yaml:"bot"`
}

func NewConfig() *Config {
	c := &Config{}

	cFile, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(cFile, c)

	if err != nil {
		log.Fatalln(err)
	}

	return c
}
