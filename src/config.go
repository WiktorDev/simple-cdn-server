package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Token string `json:"token"`
	Bind  string `json:"bind"`
}

func throwError() {
	fmt.Println("Can't start app because config.json file not found!")
	os.Exit(0)
}
func loadConfigFile() Config {
	var config Config
	jsonFile, err := os.Open("config.json")
	if err != nil {
		throwError()
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	if err := json.Unmarshal(byteValue, &config); err != nil {
		throwError()
	}
	return config
}
