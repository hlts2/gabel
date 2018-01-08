package gabel

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/hlts2/gabel/helpers"
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
		*GIO
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
	fmt.Printf(labelingMessageTmpl(g.Labels), id, text)
	for {
		in := g.StdIn.ReadLine()
		if in == "mod" {
			//TODO modify past label
		}

		sp, err := helpers.StringToIntSlice(in, ",")
		if err == nil {
			if helpers.IsContainsAllElement(g.Labels.GetValues(), sp) {
				writer.Write([]string{text, in})
				writer.Flush()
				break
			}
		}
		fmt.Print("Please re-enter:")
	}
}
