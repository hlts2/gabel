package gabel

import (
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

// OutputFileName is output file name of labeling result
var OutputFileName = "labeling_result.csv"

// Templator is text displayed template
type Templator func() string

// Gabel is core struct of gabel
type Gabel struct {
	sw     ScanWriter
	config Config
	csv    *CSV
	tmpl   template.Template
}

// NewGabel returns Gabel object
func NewGabel(sw ScanWriter, config Config, csv *CSV, templator Templator) (*Gabel, error) {
	tmpl, err := template.New(AppName).Parse(templator())
	if err != nil {
		return nil, errors.Wrap(err, "faild to parse template")
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
		g.tmpl.Execute(g.sw, g.csv.Records[i].Get())
		g.sw.WriteString(stringTables)
		g.sw.WriteString(">>> ")
		g.sw.Flush()

	Back:
		// TODO Add processing to change past labels

		s, err := g.sw.ReadLine()
		if err != nil {
			return errors.Wrap(err, "faild to read line")
		}

		labels := strings.Split(s, ",")

		if !g.config.ValidateLabels(labels) {
			g.sw.WriteString("Invlid label\n")
			g.sw.WriteString(">>> ")
			g.sw.Flush()
			goto Back
		}

		record := g.csv.Records[i]

		if record.IsAppended() {
			record.Reset()
			record.Append(labels)
		} else {
			record.Append(labels)
		}
	}

	err := g.csv.WriteCSV(OutputFileName)
	if err != nil {
		return err
	}

	return nil
}
