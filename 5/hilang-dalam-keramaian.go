package hilang_dalam_keramaian

import (
	"fmt"
	"strconv"
)

func hilangDalamKeramaian(char string) int {
	missingUnitDigitNumber, isUnitDigit := isUnitDigit(char)
	missingTensDigitNumber, isTensDigit := isTensDigit(char)
	missingHundredDigitNumber, isHundredDigit := isHundredDigit(char)

	if isUnitDigit {
		return missingUnitDigitNumber
	}

	if isTensDigit {
		return missingTensDigitNumber
	}

	if isHundredDigit {
		return missingHundredDigitNumber
	}

	fmt.Printf("\n\n Salah memasukkan data.")
	return 0
}

func isUnitDigit(str string) (int, bool) {
	isTrue := false
	isMatch := 0
	isNotMatch := 0
	missingNumber := 0

	for i := 0; i < len(str); i++ {
		if i != len(str)-1 {
			currentNumber, _ := strconv.Atoi(string(str[i]))
			nextNumber, _ := strconv.Atoi(string(str[i+1]))

			isTrue = false
			if currentNumber+1 == nextNumber {
				isTrue = true
				isMatch++
			} else {
				isNotMatch++
				if isMatch > 1 || (currentNumber+1 == nextNumber) {
					missingNumber = currentNumber + 1
				}
				if isNotMatch > 1 {
					break
				}
			}
		}
	}
	return missingNumber, isTrue
}

func isTensDigit(str string) (int, bool) {
	isTrue := false
	missingNumber := 0

	if len(str)%2 == 0 {
		isMatch := 0
		isNotMatch := 0

		for i := 0; i < len(str); i += 2 {
			currentNumberIndex := i + 2

			currentNumber, _ := strconv.Atoi(str[i:currentNumberIndex])                   // 0:2
			nextNumber, _ := strconv.Atoi(str[currentNumberIndex : currentNumberIndex+2]) // 2:4

			isTrue = false
			if currentNumber+1 == nextNumber {
				isTrue = true
				isMatch++
			} else {
				if isMatch > 1 || (currentNumber+1 == nextNumber) {
					isTrue = true
					missingNumber = currentNumber + 1
					break
				}
				isNotMatch++
				if isNotMatch > 1 {
					break
				}
			}
		}
	}
	return missingNumber, isTrue
}

func isHundredDigit(str string) (int, bool) {
	isTrue := false
	missingNumber := 0
	if len(str)%3 == 0 {
		isMatch := 0
		isNotMatch := 0
		for i := 0; i < len(str); i += 3 {
			currentNumberIndex := i + 3

			currentNumber, _ := strconv.Atoi(str[i:currentNumberIndex])                   // 0:3
			nextNumber, _ := strconv.Atoi(str[currentNumberIndex : currentNumberIndex+3]) // 3:6

			isTrue = false
			if currentNumber+1 == nextNumber {
				isTrue = true
				isMatch++
			} else {
				if isMatch > 1 || (currentNumber+1 == nextNumber) {
					isTrue = true
					missingNumber = currentNumber + 1
					break
				}
				isNotMatch++
				if isNotMatch > 1 {
					break
				}
			}
		}
	}
	return missingNumber, isTrue
}
