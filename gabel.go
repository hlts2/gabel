package gabel

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

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

//Run execute labeling process
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
again:
	fmt.Print(messageTmpl(id, text, g.Labels))
	for {
		in := g.StdIn.ReadLine()
		if in == "mod" {
			for {
				//enter past labeling ID
				n, err := strconv.Atoi(g.StdIn.ReadLine())
				if err == nil && (0 < n && n < id) {
					g.modifyOfLabeling(n, writer)
					goto again
				}
				fmt.Print("Please re-enter:")
			}
		}

		//convert comma-separated string to int slice
		//"1, 2, 3" => []int{1, 2, 3}
		isl, err := helpers.StringToIntSlice(in, ",")
		if err == nil {
			if helpers.IsContainsAllElement(g.Labels.GetValues(), isl) {
				writer.Write([]string{text, in})
				writer.Flush()
				break
			}
		}
		fmt.Print("Please re-enter:")
	}
}

func (g Gabel) modifyOfLabeling(id int, writer *csv.Writer) {
}
