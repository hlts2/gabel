package gabel

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//LabelingInfo is config.yaml Model
type LabelingInfo struct {
	Path   string          `yaml:"path"` //The file to which you want to label
	Labels `yaml:"labels"` //Details of Labeling
}

//Labels is Label slice
type Labels []Label

//Label is detail of labeling
type Label struct {
	Name  string `yaml:"name"`
	Value int    `yaml:"value"`
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

//GetValues returns Labels Value
func (labels Labels) GetValues() []int {
	s := make([]int, 0)
	for _, label := range labels {
		s = append(s, label.Value)
	}
	return s
}

//GetNames return Labels Name
func (labels Labels) GetNames() []string {
	s := make([]string, 0)
	for _, label := range labels {
		s = append(s, label.Name)
	}
	return s
}
