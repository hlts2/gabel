package helpers

import (
	"testing"
)

func TestIsContainsAllElement(t *testing.T) {
	var datas []int
	var targets []int

	datas = []int{1, 2, 3, 4, 5}
	targets = []int{2, 5}
	if !IsContainsAllElement(datas, targets) {
		t.Errorf("the result is incorect: datas = %v, targets = %v", datas, targets)
	}

	datas = []int{6, 7, 8, 9, 10}
	targets = []int{6, 20}
	if IsContainsAllElement(datas, targets) {
		t.Errorf("the result is incorect: datas = %v, targets = %v", datas, targets)
	}
}

func TestStringToIntSlice(t *testing.T) {
	nums, err := StringToIntSlice("1,2,3", ",")
	if err != nil {
		t.Errorf("StringToIntSlice() err is error: %v", err)
	}

	if len(nums) != 3 {
		t.Errorf("StringToIntSlice() nums elements is incorect: %v", nums)
	}

	nums, err = StringToIntSlice("5,A,a", ",")
	if err == nil {
		t.Errorf("StringToIntSlice() err is nil: %v", err)
	}
}
