package gabel

import (
	"encoding/csv"
	"fmt"
	"io"
)

//Output file Config for the result
const (
	DirForResult   = "GabelResult"
	OutputFileName = "labeld.csv"
)

//Gabel is gabel base struct
type Gabel struct {
	LabelingInfo
	Stdin io.Reader
}

//Run labeling process
func (c Gabel) Run(reader *csv.Reader, writer *csv.Writer) error {
	for i := 0; ; i++ {
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
