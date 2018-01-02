package gabel

//LabelingInfo is config.yaml Model
type LabelingInfo struct {
	Path   string  `yaml:"path"`
	Labels []Label `yaml:"labels"`
}

//Label is Label Field Model
type Label struct {
	Name   string `yaml:"name"`
	Values []int  `yaml:"values"`
}
