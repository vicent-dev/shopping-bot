package storage

import (
	"encoding/json"
	"github.com/en-vee/alog"
	"os"
)

var (
	cartsFilePath = os.Getenv("STORE_JSON_PATH")
)

func Load[T any](ts T) {
	bs, _ := os.ReadFile(cartsFilePath)
	err := json.Unmarshal(bs, &ts)
	if err != nil {
		alog.Error(err.Error())
	}
}

func Store[T any](data T) {
	bs, _ := json.MarshalIndent(data, "", " ")
	err := os.WriteFile(cartsFilePath, bs, 0644)

	if err != nil {
		alog.Error(err.Error())
	}
}
