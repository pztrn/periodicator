package config

import (
	"os"
	"path/filepath"
	"strings"

	"go.dev.pztrn.name/periodicator/internal/gitlab"
	"go.dev.pztrn.name/periodicator/internal/tasks"
	"gopkg.in/yaml.v2"
)

// Config is a global configuration structure.
type Config struct {
	Gitlab gitlab.Config  `yaml:"gitlab"`
	Tasks  []tasks.Config `yaml:"tasks"`
}

// Parse tries to parse configuration and returns filled structure.
func Parse() *Config {
	configPath, found := os.LookupEnv("GPT_CONFIG")
	if !found {
		panic("No configuration file path provided in 'GPT_CONFIG' environment variable!")
	}

	if strings.HasPrefix(configPath, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic("Failed to get user's home directory: " + err.Error())
		}

		configPath = strings.Replace(configPath, "~", homeDir, 1)
	}

	configPath, absErr := filepath.Abs(configPath)
	if absErr != nil {
		panic("Failed to get absolute path for '" + configPath + "': " + absErr.Error())
	}

	data, readErr := os.ReadFile(configPath)
	if readErr != nil {
		panic("Failed to read configuration file data: " + readErr.Error())
	}

	// nolint:exhaustivestruct
	c := &Config{}

	if err := yaml.Unmarshal(data, c); err != nil {
		panic("Failed to unmarshal YAML data: " + err.Error())
	}

	return c
}
