package internal

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port string   `json:"port"`
	Db   DbConfig `json:"db"`
}

type DbConfig struct {
	User   string `json:"user"`
	Passwd string `json:"passwd"`
	Addr   string `json:"addr"`
	DbName string `json:"dbName"`
}

func ParseFromFile(path string) (cfg Config, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(data, &cfg)
	return
}
