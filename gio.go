package gabel

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/hlts2/gabel/helpers"
)

type (
	//GIO is gabel io structure
	GIO struct {
		StdIn
		WFile, RFile *os.File
	}

	//StdIn is standard input
	StdIn struct {
		sc *bufio.Scanner
	}
)

//NewGIO returns GIO instance
func NewGIO(csvPath string) (*GIO, error) {
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
		return nil, err
	}

	std := StdIn{
		sc: bufio.NewScanner(os.Stdin),
	}

	g := &GIO{
		StdIn: std,
		WFile: wf,
		RFile: rf,
	}

	return g, nil
}

//Closes close read and writer files
func (g *GIO) Closes() {
	g.WFile.Close()
	g.RFile.Close()
}

//ReadLine receive input
func (s StdIn) ReadLine() string {
	s.sc.Scan()
	return s.sc.Text()
}
