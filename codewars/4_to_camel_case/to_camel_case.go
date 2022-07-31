package to_camel_case

import "strings"

// Average time complexity: O(n)
// Worst time complexity: 	O(n)
// Space complexity: 				O(1)
func ToCamelCase(str string) string {
	if str == "" {
		return str
	}

	delimiters := [2]rune{'_', '-'}
	strList := strings.FieldsFunc(str, createMatchRune(delimiters[:]))

	for index, value := range strList[1:] {
		strList[index+1] = strings.Title(value)
	}

	return strings.Join(strList, "")
}

func createMatchRune(runes []rune) func(rune) bool {
	// Average time complexity:O(log n)
	// Worst time complexity: O(n)
	// Space complexity: 			O(1)
	return func(value rune) bool {
		for _, matcher := range runes {
			if value == matcher {
				return true
			}
		}

		return false
	}
}
