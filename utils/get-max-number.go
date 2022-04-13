package utils

func GetMaxNumber(numberList []int) int {
	max := 0

	for _, num := range numberList {
		if num > max {
			max = num
		}
	}

	return max
}
