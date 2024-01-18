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
	if err = cleanenv.ReadConfig(CONFIG_PATH, &cfg); err != nil {
		return nil, fmt.Errorf("failed reading config: %w", err)
	}

	return &cfg, nil
}

func fetchConfig(flagArgs args) error {
	if _, err := os.Stat(CONFIG_PATH); os.IsNotExist(err) {
		err = writeDefaultConfig(CONFIG_PATH, flagArgs)
		if err != nil {
			return err
		}
	}

	return nil
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

	defaultConfig := fmt.Sprintf(`port: %v
read_timeout: %s
write_timeout: %s
`, port, readTimeout, writeTimeout)

	if err := os.WriteFile(filePath, []byte(defaultConfig), 0644); err != nil {
		return fmt.Errorf("error writing default config file: %w", err)
	}

	fmt.Println("default config file created:", filePath)

	return nil
}

type args struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func processArgs() (args, error) {
	var a args

	f := flag.NewFlagSet("Moonlogs", flag.ExitOnError)
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
