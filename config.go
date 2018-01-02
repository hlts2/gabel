package gabel

//Config is config.yaml Model
type Config struct {
	Path   string  `yaml:"path"`
	Labels []Label `yaml:"labels"`
}

//Label is Label Field Model
type Label struct {
	Name   string `yaml:"name"`
	Values []int  `yaml:"values"`
}
