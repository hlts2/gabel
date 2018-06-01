package gabel

import (
	"bytes"
	"encoding/csv"
	"reflect"
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

func TestAppend(t *testing.T) {
	tests := []struct {
		record   *Record
		values   []string
		expected []string
	}{
		{
			record: &Record{
				columns: []string{"a", "b"},
			},
			values:   []string{"c"},
			expected: []string{"a", "b", "c"},
		},
	}

	for _, test := range tests {
		test.record.Append(test.values)

		got := test.record.Get()

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("Append expected: %v, got %v", test.expected, got)
		}
	}
}

func TestReset(t *testing.T) {
	tests := []struct {
		record   *Record
		values   []string
		expected []string
	}{
		{
			record: &Record{
				columns:       []string{"a", "b"},
				initialLength: 2,
			},
			values:   []string{"c", "d", "e"},
			expected: []string{"a", "b"},
		},
	}

	for _, test := range tests {
		test.record.Append(test.values)

		test.record.Reset()

		got := test.record.Get()

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("Reset expected: %v, got: %v", test.expected, got)
		}
	}
}
