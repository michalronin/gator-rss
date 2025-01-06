package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	configLocation, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(configLocation)
	if err != nil {
		return Config{}, err
	}
	var result Config
	if err := json.Unmarshal(data, &result); err != nil {
		return Config{}, err
	}
	return result, nil
}

func (c *Config) SetUser(username string) {
	c.CurrentUserName = username
	if err := write(c); err != nil {
		log.Fatal(err)
	}
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + configFileName, nil
}

func write(cfg *Config) error {
	newData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	configLocation, err := getConfigFilePath()
	if err != nil {
		return err
	}
	if err := os.WriteFile(configLocation, newData, 0644); err != nil {
		return err
	}
	return nil
}
