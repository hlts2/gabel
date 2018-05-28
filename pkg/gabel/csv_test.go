package gabel

import (
	"bytes"
	"encoding/csv"
	"testing"
)

func TestNewCSV(t *testing.T) {
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
