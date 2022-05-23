package storage

import (
	"encoding/json"
	"github.com/en-vee/alog"
	"os"
)

func Load[T any](ts T) {
	bs, _ := os.ReadFile(os.Getenv("STORE_JSON_PATH"))
	err := json.Unmarshal(bs, &ts)
	if err != nil {
		alog.Error(err.Error())
	}
}

func Store[T any](data T) {
	bs, _ := json.MarshalIndent(data, "", " ")
	err := os.WriteFile(os.Getenv("STORE_JSON_PATH"), bs, 0644)

	if err != nil {
		alog.Error(err.Error())
	}
}
