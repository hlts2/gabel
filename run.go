package gabel

import "io"

//Config is Gabel Config structure
type Config struct {
	LabelingInfo
	Stdin io.Reader
}

//Run labeling process
func (c *Config) Run() error {
	return nil
}
