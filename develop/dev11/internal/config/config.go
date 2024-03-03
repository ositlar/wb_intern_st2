package config

import (
	"encoding/json"
	"io"
	"os"
)

type config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func ConfigureServer(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	var config config
	err = json.Unmarshal(content, &config)
	if err != nil {
		return "", err
	}
	addr := config.Host + ":" + config.Port
	return addr, nil
}
