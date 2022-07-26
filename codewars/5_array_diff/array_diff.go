package array_diff

func ArrayDiff(source, intersection []int) []int {
	if len(source) == 0 || len(intersection) == 0 {
		return source
	}

	intersectionMap := make(map[int]interface{}, len(intersection))
	result := []int{}

	for _, value := range intersection {
		intersectionMap[value] = nil
	}
	for _, value := range source {
		_, found := intersectionMap[value]

		if !found {
			result = append(result, value)
		}
	}

	return result
}
