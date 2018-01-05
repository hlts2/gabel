package gabel

import (
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
