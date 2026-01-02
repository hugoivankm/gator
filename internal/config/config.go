package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var homeDir = getUserHomeDir()

const configFileName = ".gatorconfig.json"

var configPath = homeDir + "/" + configFileName

func GetConfigFilePath() string {
	return configPath
}

func getUserHomeDir() string {
	userHomedir, err := os.UserHomeDir()

	if err != nil {
		fmt.Println("Unable to locate user home directory, switching to program current directory")
		currentDir, err := os.Getwd()
		if err != nil {
			return "~"
		}
		return currentDir
	}
	return userHomedir
}

func (c *Config) SetUser(currentUserName string) {
	if strings.TrimSpace(currentUserName) == "" {

	}
	c.CurrentUserName = currentUserName
}

func Read() (Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal("Error opening configuration file:", err)
		return Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatal("Error decoding json configuration:", err)
		return Config{}, err

	}
	return config, nil
}

func Write(cfg Config) error {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal("Error opening configuration file:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	err = encoder.Encode(cfg)
	if err != nil {
		fmt.Println("Error writing gator config:", err)
		return err
	}
	return nil
}

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}
