package lib_slice

import "golang.org/x/exp/constraints"

func Reverse[Type interface{}](source []Type) []Type {
	result := []Type{}
	sourceLen := len(source)

	for index := range source {
		result = append(result, source[sourceLen-index-1])
	}

	return result
}

func Max[Value constraints.Ordered](values []Value) Value {
	maxValue := values[0]

	for _, value := range values[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}
