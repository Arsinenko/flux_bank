package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Core struct {
		Address string
	}
	Analytics struct {
		Address string
	}
	Host struct {
		Address string
	}
}

func LoadConfig() (*Config, error) {
	yamlFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
