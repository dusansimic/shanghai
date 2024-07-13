package shanghai

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"github.com/dusansimic/shanghai/file"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Engine string
	File   file.File
}

type config struct {
	Engine string `yaml:"engine"`
}

func SearchConfigFile() (string, error) {
	f, err := xdg.SearchConfigFile("shanghai/config.yaml")
	if err != nil {
		return "", fmt.Errorf("failed to find config file: %w", err)
	}

	return f, nil
}

func ReadConfig(f string) (*Config, error) {
	d, err := os.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	cfglocal := &config{}
	if yaml.Unmarshal(d, &cfglocal) != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	cfg := &Config{
		Engine: cfglocal.Engine,
	}

	return cfg, nil
}
