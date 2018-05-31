package gabel

import (
	"bytes"
	"testing"
)

func TestReadLine(t *testing.T) {
	tests := []struct {
		in       *bytes.Buffer
		expected string
	}{
		{
			in:       bytes.NewBuffer([]byte("hoge\n")),
			expected: "hoge",
		},
		{
			in:       bytes.NewBuffer([]byte("hoge\nvar")),
			expected: "hoge",
		},
		{
			in:       bytes.NewBuffer([]byte("")),
			expected: "",
		},
		{
			in:       bytes.NewBuffer([]byte("\n")),
			expected: "",
		},
	}

	for _, test := range tests {
		sw := NewScanWriter(test.in, nil)
		if sw == nil {
			t.Error("NewScanWriter is nil")
		}

		got := sw.ReadLine()

		if test.expected != got {
			t.Errorf("ReadLine expected: %v, got: %v", test.expected, got)
		}
	}
}
