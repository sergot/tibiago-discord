package bot

import (
	"os"

	"github.com/sergot/tibiago/src/models"
	"gopkg.in/yaml.v3"
)

func LoadConfig(filepath string) (*models.Config, error) {
	config := models.Config{}

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
