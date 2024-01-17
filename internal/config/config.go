package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	DBPath       string        `yaml:"db_path"`
}

func Load() (*Config, error) {
	path, err := fetchConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed loading config: %w", err)
	}

	var cfg Config
	if err = cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("failed reading config: %w", err)
	}

	return &cfg, nil
}

func fetchConfigPath() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	if configPath == "" {
		defaultConfigDir, err := os.UserConfigDir()
		if err != nil {
			return "", fmt.Errorf("error getting config directory: %w", err)
		}

		configPath = filepath.Join(defaultConfigDir, "moonlogs", "config.yaml")
	}

	configDir := filepath.Dir(configPath)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return "", fmt.Errorf("error creating config file: %w", err)
		}

		fmt.Println("config directory created:", configDir)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = writeDefaultConfig(configPath)
		if err != nil {
			return "", err
		}
	}

	return configPath, nil
}

func writeDefaultConfig(filePath string) error {
	defaultConfig := []byte(`port: 4200
read_timeout: 5s
write_timeout: 1s
db_path: ./database.sqlite
`)

	if err := os.WriteFile(filePath, defaultConfig, 0644); err != nil {
		return fmt.Errorf("error writing default config file: %w", err)
	}

	fmt.Println("default config file created:", filePath)

	return nil
}
