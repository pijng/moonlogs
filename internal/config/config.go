package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	PORT          = 4200
	READ_TIMEOUT  = 5 * time.Second
	WRITE_TIMEOUT = 1 * time.Second
	DB_PATH       = "/opt/moonlogs/database.sqlite"
	CONFIG_PATH   = "/opt/moonlogs/config.yaml"
)

type Config struct {
	Port         int           `yaml:"port"`
	DBPath       string        `yaml:"db_path"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
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

	return &cfg, nil
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
read_timeout: %s
write_timeout: %s
`, flagArgs.Port, flagArgs.DBPath, flagArgs.ReadTimeout, flagArgs.WriteTimeout)

	if err := os.WriteFile(filePath, []byte(defaultConfig), 0644); err != nil {
		return fmt.Errorf("error writing default config file: %w", err)
	}

	fmt.Println("default config file created:", filePath)

	return nil
}

type args struct {
	Config       string
	Port         int
	DBPath       string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func processArgs() (args, error) {
	var a args

	f := flag.NewFlagSet("Moonlogs", flag.ExitOnError)
	f.StringVar(&a.Config, "config", CONFIG_PATH, "path to config")
	f.IntVar(&a.Port, "port", PORT, "port to run moonlogs on")
	f.StringVar(&a.DBPath, "db-path", DB_PATH, "db path to connect to")
	f.DurationVar(&a.WriteTimeout, "write-timeout", WRITE_TIMEOUT, "write timeout duration")
	f.DurationVar(&a.ReadTimeout, "read-timeout", READ_TIMEOUT, "read timeout duration")

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
