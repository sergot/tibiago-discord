package bot

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Bot struct {
		a string
	}
}

func LoadConfig(filepath string) (*Config, error) {
	config := Config{}

	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
