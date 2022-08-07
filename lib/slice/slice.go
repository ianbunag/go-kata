package lib_slice

func Reverse[Type interface{}](source []Type) []Type {
	result := []Type{}
	sourceLen := len(source)

	for index := range source {
		result = append(result, source[sourceLen-index-1])
	}

	return result
}
