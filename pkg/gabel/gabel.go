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
	scanner *bufio.Scanner
	writer  *bufio.Writer
	config  *Config
	csv     *CSV
	tmpl    *template.Template
}

// NewGabel returns Gabel object
func NewGabel(in io.Reader, out io.Writer, config *Config, csv *CSV, templator Templator) (*Gabel, error) {
	tmpl, err := template.New(AppName).Parse(templator())
	if err != nil {
		return nil, errors.Wrap(err, "faild to create object")
	}

	return &Gabel{
		scanner: bufio.NewScanner(in),
		writer:  bufio.NewWriter(out),
		config:  config,
		csv:     csv,
		tmpl:    tmpl,
	}, nil
}

// Run starts labeling
func (g *Gabel) Run(startPos, endPos int) error {
	stringTables := g.config.StringTables()

	for i := startPos; i < endPos; i++ {
		g.tmpl.Execute(g.writer, g.csv.Records[i])
		g.writer.WriteString(stringTables)
		g.writer.WriteString(">>> ")
		g.writer.Flush()

	Back:
		labels := strings.Split(",", g.scanner.Text())

		if !g.config.ValidateLabels(labels) {
			g.writer.WriteString("Invlid label\n")
			g.writer.WriteString(">>> ")
			g.writer.Flush()
			goto Back
		}

		g.csv.Records[i] = append(g.csv.Records[i], labels...)

	}

	return nil
}
