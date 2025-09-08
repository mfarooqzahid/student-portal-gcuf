package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Portal struct {
		BaseURL          string `yaml:"base_url"`
		Profile          string `yaml:"profile"`
		Login            string `yaml:"login"`
		Academics        string `yaml:"academics"`
		AcademicsDetails string `yaml:"academics_details"`
	} `yaml:"portal"`
}

func LoadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	return cfg
}
