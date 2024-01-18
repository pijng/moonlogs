package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	PORT          = 4200
	READ_TIMEOUT  = 5 * time.Second
	WRITE_TIMEOUT = 1 * time.Second
	DB_PATH       = "./database.sqlite"
)

type Config struct {
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	DBPath       string        `yaml:"db_path"`
}

func Load() (*Config, error) {
	flagArgs, err := processArgs()
	if err != nil {
		return nil, err
	}

	path, err := fetchConfigPath(flagArgs)
	if err != nil {
		return nil, fmt.Errorf("failed loading config: %w", err)
	}

	var cfg Config
	if err = cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("failed reading config: %w", err)
	}

	return &cfg, nil
}

func fetchConfigPath(flagArgs args) (string, error) {
	configPath := flagArgs.ConfigPath

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
		err = writeDefaultConfig(configPath, flagArgs)
		if err != nil {
			return "", err
		}
	}

	return configPath, nil
}

func writeDefaultConfig(filePath string, flagArgs args) error {
	port := PORT
	if flagArgs.Port > 0 {
		port = flagArgs.Port
	}

	writeTimeout := WRITE_TIMEOUT
	if flagArgs.WriteTimeout > 0 {
		writeTimeout = flagArgs.WriteTimeout
	}

	readTimeout := READ_TIMEOUT
	if flagArgs.ReadTimeout > 0 {
		readTimeout = flagArgs.ReadTimeout
	}

	dbPath := DB_PATH
	if flagArgs.DBPath != "" {
		dbPath = flagArgs.DBPath
	}

	defaultConfig := fmt.Sprintf(`port: %v
read_timeout: %s
write_timeout: %s
db_path: %s
`, port, readTimeout, writeTimeout, dbPath)

	if err := os.WriteFile(filePath, []byte(defaultConfig), 0644); err != nil {
		return fmt.Errorf("error writing default config file: %w", err)
	}

	fmt.Println("default config file created:", filePath)

	return nil
}

type args struct {
	ConfigPath   string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DBPath       string
}

func processArgs() (args, error) {
	var a args

	f := flag.NewFlagSet("Moonlogs", flag.ExitOnError)
	f.StringVar(&a.ConfigPath, "config", "", "Path to configuration file")
	f.IntVar(&a.Port, "port", 4200, "port to run moonlogs on")
	f.DurationVar(&a.WriteTimeout, "write-timeout", 1*time.Second, "write timeout duration")
	f.DurationVar(&a.ReadTimeout, "read-timeout", 5*time.Second, "read timeout duration")

	fu := f.Usage
	f.Usage = func() {
		fu()
		fmt.Fprintln(f.Output())
	}

	err := f.Parse(os.Args[1:])
	if err != nil {
		return args{}, err
	}

	return a, nil
}
