package lib

func PowerUInt(base, exponent uint) uint {
	var value uint = 1

	for ; exponent > 0; exponent -= 1 {
		value *= base
	}

	return value
}
