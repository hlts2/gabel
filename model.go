package gabel

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

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

//LoadLabelingInfoWithGivenConfigPath is load config file of given path
func LoadLabelingInfoWithGivenConfigPath(path string, l *LabelingInfo) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(b, l); err != nil {
		return err
	}

	return nil
}
