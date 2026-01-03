package config

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (c *Config) SetUser(currentUserName string) error {
	if strings.TrimSpace(currentUserName) == "" {
		return errors.New("user cannot be an empty string.")
	}
	c.CurrentUserName = currentUserName
	err := Write(*c)
	if err != nil {
		os.Exit(1)
		return fmt.Errorf("error updating configuration file")
	}
	return nil
}

func Read() (*Config, error) {
	file, err := os.Open(GetConfigFilePath())
	if err != nil {
		err = fmt.Errorf("failed to open configuration file %s: %w ", GetConfigFilePath(), err)
		return &Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		err = fmt.Errorf("error decoding json configuration: %w", err)
		return &Config{}, err

	}
	return &config, nil
}

func Write(cfg Config) error {
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("error opening configuration file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	err = encoder.Encode(cfg)
	if err != nil {
		fmt.Println("Error writing gator configuration:", err)
		return err
	}
	return nil
}

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}
