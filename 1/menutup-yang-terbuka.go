package menutup_yang_terbuka

import (
	"fmt"
	"strings"

	"github.com/maslow123/go-simple-hack/utils"
)

func menutupYangTerbuka(char string) bool {
	specialCharacter := []string{"{", "}", "(", ")", "[", "]"}
	stripString := utils.StripSpaces(char)

	stringWithoutLetter := removeLetter(specialCharacter, stripString)
	isValid := checkBracket(stringWithoutLetter)

	return isValid
}

func removeLetter(specialCharacter []string, stripString string) string {
	var stringWithoutLetter string

	for _, sc := range specialCharacter {
		specialChar := string(sc)
		for _, letter := range stripString {
			str := string(letter)

			if specialChar == str {
				stringWithoutLetter = fmt.Sprintf("%s%s", stringWithoutLetter, str)
			}
		}
	}

	return stringWithoutLetter
}

func checkBracket(stringWithoutLetter string) bool {
	for _, _ = range stringWithoutLetter {
		stringWithoutLetter = strings.Replace(stringWithoutLetter, "{}", "", -1)
		stringWithoutLetter = strings.Replace(stringWithoutLetter, "()", "", -1)
		stringWithoutLetter = strings.Replace(stringWithoutLetter, "[]", "", -1)
	}

	if len(stringWithoutLetter) > 0 {
		return false
	}
	return true
}
