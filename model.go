package gabel

//LabelingInfo is config.yaml Model
type LabelingInfo struct {
	Path   string  `yaml:"path"`   //The file to which you want to label
	Labels []Label `yaml:"labels"` //Details of Labeling
}

//Label is detail of labeling
type Label struct {
	Name   string `yaml:"name"`
	Values []int  `yaml:"values"`
}
