package helpers

import (
	"strconv"
	"strings"
)

//IsContainsAllElement checks whether it contains all elements of targets
func IsContainsAllElement(slice []int, targets []int) bool {
	t := make([]int, 0)

	for _, target := range targets {
		for _, v := range slice {
			if v == target {
				t = append(t, v)
			}
		}
	}
	return len(t) == len(targets)
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
