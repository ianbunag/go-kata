package lib_unsigned

import "golang.org/x/exp/constraints"

func Power[Value constraints.Unsigned](base, exponent Value) Value {
	var value Value = 1

	for ; exponent > 0; exponent -= 1 {
		value *= base
	}

	return value
}

func Min[Value constraints.Unsigned](x, y Value) Value {
	if x < y {
		return x
	}

	return y
}
