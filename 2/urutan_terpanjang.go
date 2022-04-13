package urutan_terpanjang

import (
	"fmt"
	"strings"

	"github.com/maslow123/go-simple-hack/utils"
)

func urutanTerpanjang(char string) int {
	sliceString := strings.Split(char, " ")
	numbers := utils.SliceStringToSliceInt(sliceString)

	numberList := groupingNumber(numbers)
	maxGroup := getMaxNumberFromGroup(numberList)

	return maxGroup
}

func groupingNumber(numbers []int) []string {
	var numberList []string

	var str string
	for i, number := range numbers {
		str = fmt.Sprintf("%s %d", str, number)

		if i+1 != len(numbers) && number+1 != numbers[i+1] {
			numberList = append(numberList, str)
			str = ""
		}
	}

	return numberList
}

func getMaxNumberFromGroup(numberList []string) int {
	maxGroup := 0
	for _, number := range numberList {
		len := len(strings.Split(number, " ")) - 1

		if len > maxGroup {
			maxGroup = len
		}
	}

	return maxGroup
}
