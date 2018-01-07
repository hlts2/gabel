package gabel

import (
	"encoding/csv"
	"io"
)

//Output file Config for the result
const (
	DirForResult   = "GabelResult"
	OutputFileName = "labeld.csv"
)

type (

	//Gabel is gabel base struct
	Gabel struct {
		LabelingInfo
		Stdin io.Reader
	}
)

//Run labeling process
func (g Gabel) Run(reader *csv.Reader, writer *csv.Writer) error {
	for i := 0; ; i++ {
		records, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if len(records) > 0 {
			g.labeling(i, records[0], writer)
		}
	}
	return nil
}

func (g Gabel) labeling(id int, text string, writer *csv.Writer) {

}
