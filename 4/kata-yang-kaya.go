package kata_yang_kaya

import (
	"fmt"
	"strings"

	"github.com/maslow123/go-simple-hack/utils"
)

func kataYangKaya(char string) int {
	sliceString := strings.Split(char, " ")
	getUniqueCharPerWord := getUniqueCharPerWord(sliceString)
	maxUniqueCharacter := utils.GetMaxNumber(getUniqueCharPerWord)

	return maxUniqueCharacter
}

func getUniqueCharPerWord(sliceString []string) []int {
	totalUniqueCharPerWord := []int{}

	for _, word := range sliceString {
		uniqueString := checkUniqueLetter(word)
		totalUniqueCharPerWord = append(totalUniqueCharPerWord, len(uniqueString))
	}

	return totalUniqueCharPerWord
}

func checkUniqueLetter(s string) string {
	uniqueString := ""
	m := make(map[rune]uint, len(s))
	for _, r := range s {
		m[r]++
	}

	noUniqueChar := []string{}

	for _, r := range s {
		if m[r] == 1 {
			uniqueString = fmt.Sprintf("%s%s", uniqueString, string(r))
		} else {
			isExists := false
			if len(noUniqueChar) < 1 {
				uniqueString = fmt.Sprintf("%s%s", uniqueString, string(r))
				noUniqueChar = append(noUniqueChar, string(r))
			}
			for _, letter := range noUniqueChar {
				if !isExists && (letter != string(r)) {
					noUniqueChar = append(noUniqueChar, string(r))
					isExists = true
					uniqueString = fmt.Sprintf("%s%s", uniqueString, string(r))
				}
			}
		}
	}
	return uniqueString
}
