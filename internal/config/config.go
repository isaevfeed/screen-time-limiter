package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Service struct {
		Env     string        `yaml:"env"`
		Host    string        `yaml:"host"`
		Port    string        `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"service"`
	Db     Database `yaml:"database"`
	DbTest Database `yaml:"database-test"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func Load(filePath string) (*Config, error) {
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %w", err)
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(configFile, cfg); err != nil {
		return nil, fmt.Errorf("yaml.Unmarshal: %w", err)
	}

	return cfg, nil
}
