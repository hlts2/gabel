package gabel

import (
	"os"
	"unsafe"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// Config is CLI config object
type Config struct {
	Path   string `yaml:"path"`
	Tables Tables `yaml:"tables"`
}

// Table is correspondence table of label
type Table struct {
	Name  string `yaml:"name"`
	Label string `yaml:"label"`
}

// Tables is Table slice
type Tables []Table

// LoadConfig loads Config object from given the path
func LoadConfig(config *Config, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "faild to open config file")
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(config)
	if err != nil {
		return errors.Wrap(err, "faild to decode cofig file")
	}

	return nil
}

// ValidateLabels validates labels
func (c *Config) ValidateLabels(labels []string) bool {
	for _, lv := range labels {
		exist := false
		for _, table := range c.Tables {
			if lv == table.Label {
				exist = true
				break
			}
		}

		if !exist {
			return false
		}
	}
	return true
}

// StringTables returns string tables. The return value format is "key: value\nkey: value"
func (c *Config) StringTables() string {
	if len(c.Tables) == 0 {
		return ""
	}

	out := make([]byte, 0, c.TablesFiledSize()+(3*len(c.Tables)))

	for _, table := range c.Tables {
		out = append(out, table.Name...)
		out = append(out, ": "...)
		out = append(out, table.Label...)
		out = append(out, "\n"...)
	}

	return *(*string)(unsafe.Pointer(&out))
}

// TablesFiledSize returns field size of all tables
func (c *Config) TablesFiledSize() (cnt int) {
	for _, table := range c.Tables {
		cnt += len(table.Name) + len(table.Label)
	}

	return cnt
}
