package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	PORT          = 4200
	READ_TIMEOUT  = 1 * time.Second
	WRITE_TIMEOUT = 10 * time.Second
	DB_PATH       = "/var/lib/moonlogs/database.sqlite"
	CONFIG_PATH   = "/etc/moonlogs/config.yaml"
	GEMINI_TOKEN  = ""

	DB_SQLITE_ADAPTER = "sqlite"
)

var config *Config

type Config struct {
	Port         int           `yaml:"port"`
	DBPath       string        `yaml:"db_path"`
	DBAdapter    string        `yaml:"db_adapter"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	GeminiToken  string        `yaml:"gemini_token"`
}

func Load() (*Config, error) {
	flagArgs, err := processArgs()
	if err != nil {
		return nil, err
	}

	err = fetchConfig(flagArgs)
	if err != nil {
		return nil, fmt.Errorf("failed loading config: %w", err)
	}

	var cfg Config
	if err = cleanenv.ReadConfig(flagArgs.Config, &cfg); err != nil {
		return nil, fmt.Errorf("failed reading config: %w", err)
	}

	config = &cfg

	return &cfg, nil
}

func Get() *Config {
	return config
}

func fetchConfig(flagArgs args) error {
	if _, err := os.Stat(flagArgs.Config); os.IsNotExist(err) {
		err = writeDefaultConfig(flagArgs.Config, flagArgs)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeDefaultConfig(filePath string, flagArgs args) error {
	defaultConfig := fmt.Sprintf(`port: %v
db_path: %s
db_adapter: %s
read_timeout: %s
write_timeout: %s
`, flagArgs.Port,
		flagArgs.DBPath,
		flagArgs.DBAdapter,
		flagArgs.ReadTimeout,
		flagArgs.WriteTimeout)

	dir := filepath.Dir(filePath)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating config dir: %w", err)
	}

	if err := os.WriteFile(filePath, []byte(defaultConfig), 0644); err != nil {
		return fmt.Errorf("error writing default config file: %w", err)
	}

	log.Println("default config file created:", filePath)

	return nil
}

type args struct {
	Config       string
	Port         int
	DBPath       string
	DBAdapter    string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	GeminiToken  string
}

func processArgs() (args, error) {
	var a args

	f := flag.NewFlagSet("Moonlogs", flag.ExitOnError)
	f.StringVar(&a.Config, "config", CONFIG_PATH, "path to config")
	f.IntVar(&a.Port, "port", PORT, "port to run moonlogs on")
	f.StringVar(&a.DBPath, "db-path", DB_PATH, "db path to connect to")
	f.StringVar(&a.DBAdapter, "db-adapter", DB_SQLITE_ADAPTER, "db adapter to connect to – 'mongodb' or 'sqlite'")
	f.DurationVar(&a.WriteTimeout, "write-timeout", WRITE_TIMEOUT, "write timeout duration")
	f.DurationVar(&a.ReadTimeout, "read-timeout", READ_TIMEOUT, "read timeout duration")
	f.StringVar(&a.GeminiToken, "gemini-token", GEMINI_TOKEN, "token to access Gemini API")

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
