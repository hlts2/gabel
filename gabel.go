package gabel

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"

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
		*Gabelio
	}

	//Gabelio is gabel io base struct
	Gabelio struct {
		Stdin        io.Reader
		WFile, RFile *os.File
	}
)

//NewGabelio returns Gabelio instance
func NewGabelio(csvPath string) (*Gabelio, error) {
	if err := helpers.Mkdir(DirForResult); err != nil {
		return nil, err
	}

	name := filepath.Join(DirForResult, OutputFileName)
	wf, err := helpers.CreateFile(name, os.O_RDWR)
	if err != nil {
		return nil, err
	}

	rf, err := helpers.OpenFile(csvPath, os.O_RDONLY)
	if err != nil {
		wf.Close()
		return nil, err
	}

	g := &Gabelio{
		Stdin: os.Stdin,
		WFile: wf,
		RFile: rf,
	}

	return g, nil
}

//FilesClose Close Read and Writer File
func (g *Gabelio) FilesClose() {
	g.WFile.Close()
	g.RFile.Close()
}

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
