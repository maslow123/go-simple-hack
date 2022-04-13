package utils

import (
	"sort"
	"strconv"
)

func SliceStringToSliceInt(str []string) []int {
	ints := make([]int, len(str))

	for i, s := range str {
		ints[i], _ = strconv.Atoi(s)
	}

	sort.Ints(ints)
	return ints
}
