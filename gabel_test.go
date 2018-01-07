package gabel

import "testing"

const (
	CSVPath = "example/example.csv"
)

func TestNewGabelio(t *testing.T) {
	gio, err := NewGabelio(CSVPath)
	if err != nil {
		t.Errorf("NewGabelio(CSVPath) err is error: %v", err)
	}

	if gio == nil {
		t.Errorf("newGabelio(CSVPath) gio is nil")
	}
}
