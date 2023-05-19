package libshg

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
)

type Config struct {
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

	cfg := &Config{}
	if yaml.Unmarshal(d, &cfg) != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return cfg, nil
}
