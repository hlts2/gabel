package gabel

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/pkg/errors"
)

// CSV is csv object
type CSV struct {
	Records []Record
}

// Record is the csv record
type Record struct {
	columns       []string
	initialLength int
}

// Get returns columns
func (r *Record) Get() []string {
	return r.columns
}

// Append appends to the column
func (r *Record) Append(values []string) {
	r.columns = append(r.columns, values...)
}

// Reset reset columns to the initial state
func (r *Record) Reset() {
	r.columns = r.columns[:r.initialLength]
}

// IsUpdated returns true if the columns has been updated
func (r *Record) IsUpdated() bool {
	return len(r.columns) != r.initialLength
}

// NewCSV returns CSV object from given the path
func NewCSV(path string) (*CSV, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "faild to open csv file")
	}
	defer f.Close()

	r := csv.NewReader(f)

	recordCnt, err := getRecordCount(r)
	if err != nil {
		return nil, err
	}
	f.Seek(0, 0)

	c := &CSV{
		Records: make([]Record, 0, recordCnt),
	}

	for {
		columns, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.Wrap(err, "faild to read csv file")
		}

		c.Records = append(c.Records, Record{
			columns:       columns,
			initialLength: len(columns),
		})
	}

	return c, nil
}

func getRecordCount(reader *csv.Reader) (cnt int, err error) {
	for {
		_, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, errors.Wrap(err, "faild to read csv file")
		}
		cnt++
	}

	return cnt, nil
}
