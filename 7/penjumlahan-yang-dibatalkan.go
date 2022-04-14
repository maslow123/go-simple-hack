package penjumlahan_yang_dibatalkan

import (
	"strconv"
	"strings"
)

func penjumlahanYangDibatalkan(char string) int {
	splitString := strings.Split(char, " ")
	sum := 0
	for i, word := range splitString {
		if word == "B" {
			previousNumber, _ := strconv.Atoi(splitString[i-1])
			sum -= previousNumber
			continue
		}
		currentNumber, _ := strconv.Atoi(word)
		sum += currentNumber
	}
	return sum

}
