package gabel

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/pkg/errors"
)

// CSV is csv object of test data
type CSV struct {
	Records []Record
}

// Record is the csv record
type Record []string

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
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.Wrap(err, "faild to read csv file")
		}

		c.Records = append(c.Records, record)
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
