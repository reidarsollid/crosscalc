package crisscross

type findProduct func(slices ...[]int) int

func CrissCrossProducts(fn findProduct, ch chan int, slices ...[]int) {
	ch <- fn(slices...)
}

func FindMax(first int, second int) int {
	if first > second {
		return first
	}
	return second
}

func productSlice(ints ...int) int {
	if containsZero(ints...) {
		return 0
	}
	l := len(ints)
	retval := ints[0]
	for i := 1; i < l; i++ {
		retval *= ints[i]
	}
	return retval
}

func containsZero(ints ...int) bool {
	for _, v := range ints {
		if v == 0 {
			return true
		}
	}
	return false
}

func productOneRow(slice []int) int {
	lent := len(slice) - 3
	max := 0
	for i := 0; i < lent; i++ {
		sum := productSlice(slice[i], slice[i+1], slice[i+2], slice[i+3])
		if max < sum {
			max = sum
		}
	}
	return max
}
func ProductRow(slice ...[]int) int {
	max := 0
	for _, v := range slice {
		if max < productOneRow(v) {
			max = productOneRow(v)
		}
	}
	return max
}

func ProductRightCross(slices ...[]int) int {
	l := len(slices) - 3
	max := 0
	for j := 0; j < l; j++ {
		for i := 0; i < l; i++ {
			sum := productSlice(slices[j][i], slices[j+1][i+1], slices[j+2][i+2], slices[j+3][i+3])
			if max < sum {
				max = sum
			}
		}
	}
	return max
}

func ProductLeftCross(slices ...[]int) int {
	l := len(slices) - 3
	max := 0
	for j := 0; j < l; j++ {
		for i := 0; i < l; i++ {
			sum := productSlice(slices[j][i+3], slices[j+1][i+2], slices[j+2][i+1], slices[j+3][i])
			if max < sum {
				max = sum
			}
		}
	}
	return max
}

func ProductDiagonal(slices ...[]int) int {
	leng := len(slices[0]) - 4
	max := 0
	for j := 0; j < leng; j++ {
		for i := range slices[j] {
			sum := productSlice(slices[j][i], slices[j+1][i], slices[j+2][i], slices[j+3][i])
			if max < sum {
				max = sum
			}
		}
	}
	return max
}
