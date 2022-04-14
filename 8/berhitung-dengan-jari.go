package berhitung_dengan_jari

import (
	"strconv"
	"strings"
)

func berhitungDenganJari(char string) string {
	fingers := []string{"jempol", "telunjuk", "tengah", "manis", "kelingking"}
	splitString := strings.Split(char, " ")

	if len(splitString) > 2 {
		return "invalid-data"
	}

	start, _ := strconv.Atoi(splitString[0])
	end, _ := strconv.Atoi(splitString[1])
	index := 0
	isReverse := false

	for i := start - 1; i <= end+1; i++ {
		if index == len(fingers) || isReverse {
			isReverse = true
			if index == len(fingers) {
			}
			index--
			if index == 0 {
				isReverse = false
			}
			continue
		}

		if index == 0 || !isReverse {
			isReverse = false
			index++
		}

	}
	return fingers[index-1]
}
