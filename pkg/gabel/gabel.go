package gabel

import (
	"strconv"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

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
	Back:
		g.sw.WriteString("\nNo." + strconv.Itoa(i) + "\n")
		g.tmpl.Execute(g.sw, g.csv.Records[i].Get())
		g.sw.WriteString(stringTables)
		g.sw.WriteString(">>> ")
		g.sw.Flush()

	ReEnter:
		s, err := g.sw.ReadLine()
		if err != nil {
			return errors.Wrap(err, "faild to read line")
		}

		labels := strings.Split(s, ",")

		if g.config.IsModificationLabel(labels[0]) {
			n, err := g.getRecordNumber(i)
			if err != nil {
				return err
			}

			if err := g.Run(n, n+1); err != nil {
				return err
			}

			goto Back
		} else if !g.config.ValidateLabels(labels) {
			g.sw.WriteString("Invlid label\n")
			g.sw.WriteString(">>> ")
			g.sw.Flush()
			goto ReEnter
		}

		record := g.csv.Records[i]

		if record.IsAppended() {
			record.Reset()
			record.Append(labels)
		} else {
			record.Append(labels)
		}
	}

	return nil
}

// GenerateCSV create csv file
func (g *Gabel) GenerateCSV(path string) error {
	if err := g.csv.WriteCSV(path); err != nil {
		return errors.Wrap(err, "faild to generate csv file")
	}
	return nil
}

func (g *Gabel) getRecordNumber(endPos int) (int, error) {
	for {

	ReEnter:
		g.sw.WriteString(">>> ")
		g.sw.Flush()

		s, err := g.sw.ReadLine()
		if err != nil {
			return 0, errors.Wrap(err, "faild to read line")
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			g.sw.WriteString("Invalid number\n")
			g.sw.Flush()
			goto ReEnter
		}

		if n < 0 || n > endPos {
			g.sw.WriteString("Invalid number\n")
			g.sw.Flush()
			goto ReEnter
		}

		return n, nil
	}
}
