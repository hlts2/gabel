package helpers

import (
	"testing"
)

func TestIsContainsAllElement(t *testing.T) {
	var datas1 []int
	var datas2 []int

	datas1 = []int{1, 2, 3, 4, 5}
	datas2 = []int{2, 5}
	if !IsContainsAllElement(datas1, datas2) {
		t.Errorf("the result is incorect: datas = %v, targets = %v", datas1, datas2)
	}

	datas1 = []int{6, 7, 8, 9, 10}
	datas2 = []int{6, 200}
	if IsContainsAllElement(datas1, datas2) {
		t.Errorf("the result is incorect: datas = %v, targets = %v", datas1, datas2)
	}
}

func TestStringToIntSlice(t *testing.T) {
	var datas1 string

	datas1 = "1,2,3"
	_, err := StringToIntSlice(datas1, ",")
	if err != nil {
		t.Errorf("StringToIntSlice() err is error: %v", err)
	}

	datas1 = "5,A,a"
	_, err = StringToIntSlice(datas1, ",")
	if err == nil {
		t.Errorf("StringToIntSlice() err is nil: %v", err)
	}
}
