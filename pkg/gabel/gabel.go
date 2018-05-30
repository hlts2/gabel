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
	config    *Config
	csv       *CSV
	templator Templator
}

// NewGabel returns Gabel object
func NewGabel(config *Config, csv *CSV, templator Templator) *Gabel {
	return &Gabel{
		config:    config,
		csv:       csv,
		templator: templator,
	}
}

// Run starts labeling
func (g *Gabel) Run(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)

	tmpl, err := template.New("gabel").Parse(g.templator())
	if err != nil {
		return err
	}

	stringTables := g.config.StringTables()

	for i, record := range g.csv.Records {

		tmpl.Execute(writer, record)
		writer.WriteString(stringTables)
		writer.WriteString(">>> ")
		writer.Flush()

	Back:
		labels := strings.Split(",", scanner.Text())

		if !g.config.ValidateLabels(labels) {
			writer.WriteString("Invlid label\n")
			writer.WriteString(">>> ")
			writer.Flush()
			goto Back
		}

		g.csv.Records[i] = append(g.csv.Records[i], labels...)
	}

	return nil
}
