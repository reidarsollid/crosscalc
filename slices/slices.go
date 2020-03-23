package slices

import (
	"fmt"
	"strconv"
	"strings"
)

func TextToIntSlices(text string) [][]int {
	lines := strings.Split(text, "\n")
	slices := make([][]int, len(lines))
	for i, v := range lines {
		ssLice := strings.Split(v, " ")
		intSlice := toIntSlice(ssLice)
		slices[i] = intSlice
	}
	return slices
}

func toIntSlice(slice []string) []int {
	retval := make([]int, 0)
	for i := range slice {
		i, err := strconv.Atoi(slice[i])
		if err != nil {
			fmt.Println(err)
		}
		retval = append(retval, int(i))
	}
	return retval
}

