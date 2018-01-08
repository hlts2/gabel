package gabel

import (
	"os"
	"testing"
)

const CSVPath = "./example/example.csv"

func TestNewGIO(t *testing.T) {
	defer os.RemoveAll(DirForResult)

	gio, err := NewGIO(CSVPath)
	if err != nil {
		t.Errorf("NewGIO() err is err: %v", err)
	}

	if gio == nil {
		t.Errorf("NewGIO() gio is nil")
	}
}
