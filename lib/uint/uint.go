package lib_uint

func Power(base, exponent uint) uint {
	var value uint = 1

	for ; exponent > 0; exponent -= 1 {
		value *= base
	}

	return value
}

func Min(x, y uint) uint {
	if x < y {
		return x
	}

	return y
}
