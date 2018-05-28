package gabel

import (
	"bytes"
	"encoding/csv"
	"testing"
)

func TestNewCSV(t *testing.T) {
	path := "../../example/example.csv"

	c, err := NewCSV(path)
	if err != nil {
		t.Errorf("NewCSV is error: %v", err)
	}

	if c == nil {
		t.Errorf("NewCSV c is nil")
	}
}

func TestGetRecordCount(t *testing.T) {
	tests := []struct {
		reader   *csv.Reader
		isError  bool
		expected int
	}{
		{
			reader:   csv.NewReader(bytes.NewBuffer([]byte("hoge\nfoo\nvar"))),
			isError:  false,
			expected: 3,
		},
	}

	for i, test := range tests {
		count, err := getRecordCount(test.reader)

		isError := !(err == nil)

		if test.isError != isError {
			t.Errorf("i = %d getRecordCount expected isError: %v, got: %v", i, test.isError, isError)
		}

		if test.expected != count {
			t.Errorf("i = %d getRecordCount count expected: %v, got: %v", i, test.isError, isError)
		}
	}
}
