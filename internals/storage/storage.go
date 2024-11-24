package storage

import (
	"encoding/json"
	"log"
	"os"
)

var configFile *ConfigFileStruct

type ConfigFileStruct struct {
	NumberOfRecords int `json:"number_of_records"`
}

func GetNumberOfRecords() int {
	return configFile.NumberOfRecords
}

func SetNumberOfRecords(n int) {
	configFile.NumberOfRecords = n
	err := Save()
	if err != nil {
		log.Println(err.Error())
	}
}

func Save() error {
	content, err := json.Marshal(configFile)
	if err != nil {
		return err
	}
	err = os.WriteFile("config.json", content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadConfigFile() error {
	configFile = &ConfigFileStruct{
		NumberOfRecords: 0,
	}
	var file *os.File
	file, err := os.Open("config.json")
	if err != nil {
		file, err = os.Create("config.json")
		defer file.Close()
		Save()
		if err != nil {
			return err
		}
	}
	content, err := os.ReadFile("config.json")
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.Unmarshal(content, configFile)
	if err != nil {
		return err
	}
	return nil
}
