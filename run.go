package gabel

import (
	"encoding/csv"
	"io"
)

//Config is gabel config structure
type Config struct {
	LabelingInfo
	Stdin io.Reader
}

//Output file Config for the result
const (
	DirForResult   = "GabelResult"
	OutputFileName = "labeld.csv"
)

//Run labeling process
func (c *Config) Run(reader *csv.Reader, writer *csv.Writer) error {
	return nil
}
