package gabel

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config is CLI config object
type Config struct {
	Path   string `yaml:"path"`
	Tables Tables `yaml:"tables"`
}

// Table is correspondence table of label
type Table struct {
	Name   string   `yaml:"name"`
	Labels []string `yaml:"labels"`
}

// Tables is Table slice
type Tables []Table

// LoadConfig loads Config object from given the path
func LoadConfig(config *Config, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		return err
	}

	return nil
}

// ValidateLabels validates labels
func (c *Config) ValidateLabels(labels []string) bool {
	return false
}
