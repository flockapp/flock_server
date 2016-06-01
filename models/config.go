package models

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Config struct {
	DatabasePath string `json:"dbPath"`
	Port         string `json:"port"`
	SigningKey   string `json:"signingKey"`
}

var Conf Config

func init() {
	jsonData, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(fmt.Sprintf("Error reading config vars %v\n", err))
	}
	err = json.Unmarshal(jsonData, &Conf)
	if err != nil {
		panic("Error parsing JSON")
	}
}