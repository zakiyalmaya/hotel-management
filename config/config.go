package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type DatabaseConfig struct {
	File string `json:"file"`
}

func Init() {
	// Initialize config with default values
	config := Config{
		Server: ServerConfig{
			Port: ":3000",
		},
		Database: DatabaseConfig{
			File: "./hotel.db",
		},
	}

	// Convert config struct to JSON
	configFile, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling config:", err)
		os.Exit(1)
	}

	// Write JSON to file
	err = os.WriteFile("config.json", configFile, 0644)
	if err != nil {
		fmt.Println("Error writing config file:", err)
		os.Exit(1)
	}

	fmt.Println("Config file generated successfully as config.json")
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the JSON into the struct
	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
