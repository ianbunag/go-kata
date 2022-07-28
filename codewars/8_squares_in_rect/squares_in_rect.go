package squares_in_rect

func SquaresInRect(length int, width int) (result []int) {
	if length == width {
		return
	}

	for length != width {
		if length < width {
			length, width = width, length
		}

		result = append(result, width)
		length = length - width
	}

	result = append(result, width)
	return
}
