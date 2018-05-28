package gabel

import (
	"encoding/csv"
	"io"
	"os"
)

// CSV is csv object of test data
type CSV struct {
	Records [][]string
}

// NewCSV returns CSV object from given the path
func NewCSV(path string) (*CSV, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)

	recordCnt, err := getRecordCount(r)
	if err != nil {
		return nil, err
	}

	c := &CSV{
		Records: make([][]string, 0, recordCnt),
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
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
			return 0, err
		}
		cnt++
	}

	return cnt, nil
}
