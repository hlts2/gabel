package helpers

import (
	"strconv"
	"strings"
)

//IsContainsAllElement checks whether it contains all elements of targets
func IsContainsAllElement(ns []int, targets []int) bool {
	nsTmp := ExcluDuplicatesElementOfSlice(ns)
	t := make([]int, 0)

	for _, target := range targets {
		for _, n := range nsTmp {
			if n == target {
				t = append(t, n)
			}
		}
	}
	return len(t) == len(targets)
}

//ExcluDuplicatesElementOfSlice returns slice without duplicate elements
func ExcluDuplicatesElementOfSlice(ns []int) []int {
	t := make([]int, 0)

	for _, n := range ns {
		if !IsContains(n, t) {
			t = append(t, n)
		}
	}
	return t
}

//IsContains check if n is included in the slice
func IsContains(n int, ns []int) bool {
	for _, v := range ns {
		if n == v {
			return true
		}
	}
	return false
}

//StringToIntSlice convert from string to int slice
func StringToIntSlice(target string, delim string) ([]int, error) {
	sp := strings.Split(target, delim)

	nums := make([]int, 0)
	for _, v := range sp {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		nums = append(nums, i)
	}
	return nums, nil
}
