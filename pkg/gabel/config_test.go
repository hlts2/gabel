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

func TestValidateLabels(t *testing.T) {
	tests := []struct {
		expected bool
		labels   []string
		config   *Config
	}{
		{
			expected: true,
			labels:   []string{"1"},
			config: &Config{
				Tables: Tables{
					{
						Name:   "dog",
						Labels: []string{"1"},
					},
					{
						Name:   "cat",
						Labels: []string{"2"},
					},
				},
			},
		},
	}

	for _, test := range tests {
		got := test.config.ValidateLabels(test.labels)

		if test.expected != got {
			t.Errorf("ValidateLabels expected: %v, got: %v", test.expected, got)
		}
	}
}
