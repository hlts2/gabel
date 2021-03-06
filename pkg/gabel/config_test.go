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
				Name:             "fail",
				ModificationFlag: false,
				Label:            "0",
			},
			{
				Name:             "success",
				ModificationFlag: false,
				Label:            "1",
			},
			{
				Name:             "modify past label",
				ModificationFlag: true,
				Label:            "-1",
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
						Name:  "dog",
						Label: "1",
					},
					{
						Name:  "cat",
						Label: "2",
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

func TestStringTables(t *testing.T) {
	tests := []struct {
		expected string
		config   *Config
	}{
		{
			expected: "[1] dog\n[2] cat\n",
			config: &Config{
				Tables: Tables{
					{
						Name:  "dog",
						Label: "1",
					},
					{
						Name:  "cat",
						Label: "2",
					},
				},
			},
		},
		{
			expected: "",
			config: &Config{
				Tables: Tables{},
			},
		},
	}

	for i, test := range tests {
		got := test.config.StringTables()

		if test.expected != got {
			t.Errorf("i = %d StringTables expected: %v, got: %v", i, test.expected, got)
		}
	}
}

func TestTableFieldSize(t *testing.T) {
	tests := []struct {
		expected int
		config   *Config
	}{
		{
			expected: 8,
			config: &Config{
				Tables: Tables{
					{
						Name:  "dog",
						Label: "1",
					},
					{
						Name:  "cat",
						Label: "2",
					},
				},
			},
		},
	}

	for _, test := range tests {
		got := test.config.TablesFiledSize()

		if test.expected != got {
			t.Errorf("TablesFiledSize expected: %v, got: %v", test.expected, got)
		}
	}

}

func TestIsModificationLabel(t *testing.T) {
	tests := []struct {
		label    string
		config   *Config
		expected bool
	}{
		{
			label: "-1",
			config: &Config{
				Tables: Tables{
					{
						Name:  "dog",
						Label: "1",
					},
					{
						Name:  "cat",
						Label: "2",
					},
					{
						Name:             "modify past label",
						ModificationFlag: true,
						Label:            "-1",
					},
				},
			},
			expected: true,
		},
	}

	for _, test := range tests {
		got := test.config.IsModificationLabel(test.label)

		if test.expected != got {
			t.Errorf("IsModificationLabel expected: %v, got: %v", test.expected, got)
		}
	}
}
