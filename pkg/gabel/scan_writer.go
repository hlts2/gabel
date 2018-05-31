package gabel

import (
	"bufio"
	"io"
)

// ScanWriter is the interface that wraps the basic Read method.
type ScanWriter interface {
	ReadLine() string
	Write(p []byte) (int, error)
	WriteString(s string) (int, error)
	Flush() error
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

// Write writes the contents of p into the buffer.
func (sw *scanWriter) Write(p []byte) (int, error) {
	return sw.writer.Write(p)
}

// WriteString writes a string
func (sw *scanWriter) WriteString(s string) (int, error) {
	return sw.writer.WriteString(s)
}

// Flush writes any buffered data to the underlying io.Writer.
func (sw *scanWriter) Flush() error {
	return sw.writer.Flush()
}
