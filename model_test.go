package gabel

import (
	"reflect"
	"testing"
)

const ConfigPath = "./example/config.yaml"

func TestLoadLabelingInfoWithGivenConfigPath(t *testing.T) {
	var l LabelingInfo
	err := LoadLabelingInfoWithGivenConfigPath(ConfigPath, &l)
	if err != nil {
		t.Errorf("LoadLabelingInfoWithGivenConfigPath(path, LabelingInfo) is error: %v", err)
	}
}

func TestGetValues(t *testing.T) {
	var data1 Labels
	var data2 []int

	data1 = Labels{
		Label{Name: "notextis", Value: 0},
		Label{Name: "exist", Value: 1},
	}
	data2 = []int{0, 1}

	if !reflect.DeepEqual(data1.GetValues(), data2) {
		t.Errorf("the result is incorect: data = %v, result = %v", data1, data2)
	}

	data1 = Labels{
		Label{Name: "cat", Value: 1},
		Label{Name: "dog", Value: 2},
		Label{Name: "human", Value: 3},
	}
	data2 = []int{1, 2, 3}

	if !reflect.DeepEqual(data1.GetValues(), data2) {
		t.Errorf("the result is incorect: data = %v, result = %v", data1, data2)
	}
}

func TestGetNames(t *testing.T) {
	var data1 Labels
	var data2 []string

	data1 = Labels{
		Label{Name: "notextis", Value: 0},
		Label{Name: "exist", Value: 1},
	}
	data2 = []string{"notextis", "exist"}

	if !reflect.DeepEqual(data1.GetNames(), data2) {
		t.Errorf("the result is incorect: data = %v, result = %v", data1, data2)
	}
}
