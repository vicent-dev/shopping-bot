package config

import (
	"embed"
)

//go:embed config.yaml
var cFile embed.FS

func GetConfigFile() []byte {
	bs, _ := cFile.ReadFile("config.yaml")
	return bs
}
