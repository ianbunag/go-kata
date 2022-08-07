package lib_generate

import "strings"

const ALL_LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateString(seed string, length int) string {
	seedLength := len([]rune(seed))
	var builder strings.Builder

	wholeLength := length / seedLength
	partialLength := length % seedLength

	for i := 0; i < wholeLength; i++ {
		builder.WriteString(seed)
	}
	builder.WriteString(string([]rune(seed)[:partialLength]))

	return builder.String()
}
