package config

import (
	"encoding/json"
	utility "moneysaverapi"
	"os"
)

func getJsonFilePath() string {
	path, err := os.Getwd()
	if err != nil {
	}
	return path + "\\config.json"
}

func GetConfigDetails() AppConfiguration {
	config := AppConfiguration{}
	file, err := os.Open(getJsonFilePath())
	if err != nil {
		utility.WriteToLog(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config
}
