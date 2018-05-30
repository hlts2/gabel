package gabel

import (
	"bufio"
	"io"
	"strings"
	"text/template"
)

// Templator is text displayed template
type Templator func() string

// Gabel is core struct of gabel
type Gabel struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
	config  *Config
	csv     *CSV
}

// NewGabel returns Gabel object
func NewGabel(in io.Reader, out io.Writer, config *Config, csv *CSV) *Gabel {
	return &Gabel{
		scanner: bufio.NewScanner(in),
		writer:  bufio.NewWriter(out),
		config:  config,
		csv:     csv,
	}
}

// Run starts labeling
func (g *Gabel) Run(startPos, endPos int, templator Templator) error {
	tmpl, err := template.New(AppName).Parse(templator())
	if err != nil {
		return err
	}

	stringTables := g.config.StringTables()

	for i := startPos; i < endPos; i++ {
		tmpl.Execute(g.writer, g.csv.Records[i])
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
