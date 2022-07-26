package to_camel_case

import "strings"

func createMatchRune(runes []rune) func(rune) bool {
	return func(value rune) bool {
		for _, matcher := range runes {
			if value == matcher {
				return true
			}
		}

		return false
	}
}

func ToCamelCase(str string) string {
	if str == "" {
		return ""
	}
	delimiters := [2]rune{'_', '-'}
	strList := strings.FieldsFunc(str, createMatchRune(delimiters[:]))

	for index, value := range strList[1:] {
		strList[index+1] = strings.Title(value)
	}

	return strings.Join(strList, "")
}
