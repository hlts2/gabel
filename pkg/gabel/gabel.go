package gabel

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
func (g *Gabel) Run() error {
	return nil
}
