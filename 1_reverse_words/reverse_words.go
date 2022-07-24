package reverse_words

import "strings"

func ReverseWord(str string) string {
	runes := []byte{}

	for ctr := len(str) - 1; len(runes) < len(str); ctr -= 1 {
		runes = append(runes, str[ctr])
	}

	return string(runes)
}

func ReverseWords(str string) string {
	words := strings.Split(str, " ")

	for key, word := range words {
		words[key] = ReverseWord(word)
	}

	return strings.Join(words, " ")
}
