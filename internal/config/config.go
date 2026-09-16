package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DataBaseURL     string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	return writeToConfigFile(c)
}

func Read() (Config, error) {

	config_file_path, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error getting config file path: %v", err)
	}
	data, err := os.ReadFile(config_file_path)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %v", err)
	}

	var conf Config
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}

func getConfigFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %v", err)
	}

	return filepath.Join(home_dir, configFileName), nil
}

func writeToConfigFile(cfg Config) error {
	config_file_path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path: %v", err)
	}

	config_bytes, _ := json.Marshal(cfg)

	err = os.WriteFile(config_file_path, config_bytes, 0644)
	if err != nil {
		return fmt.Errorf("error writting to file: %v", err)
	}

	return nil
}
