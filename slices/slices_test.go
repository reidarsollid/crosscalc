package slices

import (
	"github.com/magiconair/properties/assert"
	"strconv"
	"testing"
)

var txt string = "08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08\n" +
	"49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00\n" +
	"81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65\n" +
	"52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91\n" +
	"22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80\n" +
	"24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50\n" +
	"32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70\n" +
	"01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48"

var intSlices = [][]int{
	{8, 2, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8},
	{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 0},
	{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 03, 49, 13, 36, 65},
	{52, 70, 95, 23, 4, 60, 11, 42, 69, 24, 68, 56, 01, 32, 56, 71, 37, 2, 36, 91},
	{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80},
	{24, 47, 32, 60, 99, 3, 45, 2, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50},
	{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70},
	{1, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 1, 89, 19, 67, 48},
}

func TestTextToIntSlices(t *testing.T) {
	actual := TextToIntSlices(txt)
	assert.Equal(t, len(actual), len(intSlices))
	for i, v := range intSlices {
		assert.Equal(t, v, actual[i], "Line ", strconv.Itoa(i+1))
	}
}