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
		Path: "example.csv",
		Tables: Tables{
			{
				Name:   "fail",
				Labels: []string{"0"},
			},
			{
				Name:   "success",
				Labels: []string{"1"},
			},
		},
	}

	if !reflect.DeepEqual(expected, config) {
		t.Errorf("LoadConfig config expected: %v, got: %v", expected, config)
	}
}
