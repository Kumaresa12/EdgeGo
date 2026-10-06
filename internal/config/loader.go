package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func Load(path string) (*Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	cfg := &Config{}

	err = yaml.Unmarshal(data, cfg)

	if err != nil {
		return nil, fmt.Errorf("parse yaml: %w", err)
	}

	return cfg, nil
}
