package percakapan_rahasia

import (
	"fmt"
	"strings"
)

func percakapanRahasia(char string) string {
	splitString := strings.Split(char, " ")
	word := reverseString(splitString)

	return word
}

func reverseString(str []string) string {
	var char string
	for _, word := range str {
		byte_str := []rune(word)

		for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
			byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
		}

		if char == "" {
			char = fmt.Sprintf("%s%s", char, string(byte_str))
		} else {
			char = fmt.Sprintf("%s %s", char, string(byte_str))
		}
	}

	return char
}
