package lampu_yang_hidup_mati

import (
	"strconv"
	"strings"
)

func lampuYangHidupMati(char string) int {
	splitString := strings.Split(char, " ")
	lamp, _ := strconv.Atoi(splitString[0])
	press, _ := strconv.Atoi(splitString[1])

	lampList := make([]bool, lamp)
	for i := 0; i < lamp; i++ {
		lampList[i] = false
	}

	currentLampList := pressTheSwitch(press, lampList)
	totalActiveLamp := totalActiveLamp(currentLampList)

	return totalActiveLamp
}

func pressTheSwitch(press int, lampList []bool) []bool {
	for i := 1; i <= press; i++ {
		for j := range lampList {
			if i == 0 {
				lampList[j] = true
				continue
			}
			if j+1 >= i && (j+1)%i == 0 {
				lampList[j] = !lampList[j]
			}
		}
	}

	return lampList
}

func totalActiveLamp(currentLampList []bool) int {
	var activeLamp []bool
	for _, lamp := range currentLampList {
		if lamp {
			activeLamp = append(activeLamp, lamp)
		}
	}
	return len(activeLamp)
}
