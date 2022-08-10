package dari_kursi_jadi_bui

import (
	"fmt"
	"strings"
)

func DariKursiJadiBui(data string) int {
	s := strings.Split(data, " ")

	firstWord := s[0]
	secondWord := s[1]
	secondWordFilter := secondWord

	stringFilter, counter := getMissingWord(firstWord, secondWordFilter)

	// Add unavailable word from stringFilter
	if secondWord != stringFilter {
		counter = addUnvailableWord(secondWord, stringFilter, counter)
	}

	return counter
}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}

func getMissingWord(firstWord, secondWordFilter string) (string, int) {
	counter := 0
	stringFilter := ""
	// Filter missing word from comparison
	for _, ch1 := range firstWord {
		isMatch := false
		for i, ch2 := range secondWordFilter {
			if ch1 == ch2 {
				secondWordFilter = string(delChar([]rune(secondWordFilter), i))
				stringFilter = fmt.Sprintf("%s%s", stringFilter, string(ch2))
				isMatch = true
				break
			}
		}

		if !isMatch {
			counter++
		}
	}

	return stringFilter, counter
}

func addUnvailableWord(secondWord, stringFilter string, counter int) int {
	for _, ch1 := range secondWord {
		isMatch := false

		for i, ch2 := range stringFilter {
			if ch1 == ch2 {
				stringFilter = string(delChar([]rune(stringFilter), i))
				isMatch = true
				break
			}
		}

		if !isMatch {
			counter++
		}
	}

	return counter
}
