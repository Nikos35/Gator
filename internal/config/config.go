package config

import (
	"os"
)

type Config struct {
	DataBaseUrl string `db_url`
}

func Read() Config {
	dir, _ := os.UserHomeDir()
	return Config {DataBaseUrl: dir}
	// data, _ := os.ReadFile		
}
