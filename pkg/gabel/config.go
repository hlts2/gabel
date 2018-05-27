package gabel

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config is CLI config object
type Config struct {
	Target string `yaml:"target"`
	Labels Labels `yaml:"labels"`
}

// Label is teacher label object
type Label struct {
	Name   string   `yaml:"name"`
	Values []string `yaml:"values"`
}

// Labels is Label slice
type Labels []Label

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
