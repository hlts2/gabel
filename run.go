package gabel

import (
	"encoding/csv"
	"fmt"
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
func (c Config) Run(reader *csv.Reader, writer *csv.Writer) error {
	for {
		records, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if len(records) > 0 {
			fmt.Println(records)
		}
	}
	return nil
}
