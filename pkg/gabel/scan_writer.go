package gabel

import (
	"bufio"
	"io"
)

// ScanWriter is the interface that wraps the basic Read method.
type ScanWriter interface {
	ReadLine() string
}

// scanWriter stores pointers to Scanner and a Writer.
type scanWriter struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

// NewScanWriter returns ScanWriter
func NewScanWriter(in io.Reader, out io.Writer) ScanWriter {
	return &scanWriter{
		scanner: bufio.NewScanner(in),
		writer:  bufio.NewWriter(out),
	}
}

// ReadLine returns line
func (sw *scanWriter) ReadLine() string {
	sw.scanner.Scan()
	return sw.scanner.Text()
}
