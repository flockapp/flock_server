package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	DatabasePath string `json:"dbPath"`
	Port         string `json:"port"`
	SigningKey   string `json:"signingKey"`
	GoogleApiKey string `json:"googleApi"`
}

var Conf Config

func init() {
	jsonData, err := ioutil.ReadFile("config.json")
	fmt.Printf("DBPATH: %v\n", Conf.DatabasePath)
	if err != nil {
		panic(fmt.Sprintf("Error reading config vars %v\n", err))
	}
	err = json.Unmarshal(jsonData, &Conf)
	if err != nil {
		panic("Error parsing JSON")
	}
}
