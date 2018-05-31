package gabel

import (
	"bufio"
	"io"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

// Templator is text displayed template
type Templator func() string

// Gabel is core struct of gabel
type Gabel struct {
	sw     *ScanWriter
	config Config
	csv    CSV
	tmpl   template.Template
}

// ScanWriter stores pointers to Scanner and a Writer.
type ScanWriter struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

// NewScanWriter returns ScanWriter object
func NewScanWriter(in io.Reader, out io.Writer) *ScanWriter {
	return &ScanWriter{
		scanner: bufio.NewScanner(in),
		writer:  bufio.NewWriter(out),
	}
}

// ReadLine returns line
func (sw *ScanWriter) ReadLine() string {
	sw.scanner.Scan()
	return sw.scanner.Text()
}

// NewGabel returns Gabel object
func NewGabel(sw *ScanWriter, config Config, csv CSV, templator Templator) (*Gabel, error) {
	tmpl, err := template.New(AppName).Parse(templator())
	if err != nil {
		return nil, errors.Wrap(err, "faild to NewGabel")
	}

	return &Gabel{
		sw:     sw,
		config: config,
		csv:    csv,
		tmpl:   *tmpl,
	}, nil
}

// Run starts labeling
func (g *Gabel) Run(startPos, endPos int) error {
	stringTables := g.config.StringTables()

	for i := startPos; i < endPos; i++ {
		g.tmpl.Execute(g.sw.writer, g.csv.Records[i])
		g.sw.writer.WriteString(stringTables)
		g.sw.writer.WriteString(">>> ")
		g.sw.writer.Flush()

	Back:
		labels := strings.Split(",", g.sw.ReadLine())

		if !g.config.ValidateLabels(labels) {
			g.sw.writer.WriteString("Invlid label\n")
			g.sw.writer.WriteString(">>> ")
			g.sw.writer.Flush()
			goto Back
		}

		g.csv.Records[i] = append(g.csv.Records[i], labels...)
	}

	return nil
}
