package config

import (
	"errors"
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App        App        `yaml:"app"`
	PostgresDB PostgresDB `yaml:"postgresDB"`
}
type App struct {
	Port string `yaml:"port"`
}
type PostgresDB struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
	Username string `yaml:"username"`
	DBName   string `yaml:"dbname"`
}

func NewConfig() (*Config, error) {
	path := fetchConfigPath()

	if path == "" {
		return nil, errors.New("empty config path")
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("os.ReadFile() error")
	}

	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, errors.New("yaml.Unmarshal(): " + err.Error())
	}

	return &cfg, nil
}

func fetchConfigPath() string {
	var path string

	flag.StringVar(&path, "cfgpath", "", "config path")
	flag.Parse()

	if path == "" {
		path = os.Getenv("LIR_CONFIG_PATH")
	}

	return path
}
