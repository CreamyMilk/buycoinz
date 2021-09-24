package utils

import "strings"

func ToSentenceCase(word string) string {
	if len(word) < 1 {
		return strings.ToUpper(word)
	}
	var result string
	for _, sections := range strings.Split(word, " ") {
		first := strings.ToUpper(string(sections[0]))
		result += first + strings.ToLower(sections[1:]) + " "
	}
	return result[:len(result)-1]
}
