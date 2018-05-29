package gabel

import (
	"bufio"
	"io"
	"strings"
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

	stringTables := g.config.StringTables()

	for i := range g.csv.Records {
	Back:

		// TODO output record
		writer.WriteString("")
		writer.WriteString(stringTables)
		writer.WriteString(">>> ")
		writer.Flush()

		labels := strings.Split(",", scanner.Text())

		if !g.config.ValidateLabels(labels) {
			writer.WriteString("Invlid label")
			writer.Flush()
			goto Back
		}

		g.csv.Records[i] = append(g.csv.Records[i], labels...)
	}

	return nil
}
