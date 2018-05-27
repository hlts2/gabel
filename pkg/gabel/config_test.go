package gabel

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	path := "../../example/config.yaml"

	var config Config

	err := LoadConfig(&config, path)
	if err != nil {
		t.Errorf("LoadConfig is error: %v", err)
	}

	expected := Config{
		Target: "example.csv",
		Labels: Labels{
			{
				Name:   "fail",
				Values: []string{"0"},
			},
			{
				Name:   "success",
				Values: []string{"1"},
			},
		},
	}

	if !reflect.DeepEqual(expected, config) {
		t.Errorf("LoadConfig config expected: %v, got: %v", expected, config)
	}
}
