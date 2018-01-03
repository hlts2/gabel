package gabel

import (
	"io"
	"os"
)

//Config is Gabel Config structure
type Config struct {
	LabelingInfo
	Stdin        io.Reader
	RFile, WFile *os.File
}

//Run labeling process
func (c *Config) Run() error {
	return nil
}
