package reverse_integer

const maxInt32 int = 0x7FFFFFFF         // 2147483647
const maxLeftDigits int = maxInt32 / 10 // 214748364
const maxLastDigit int = maxInt32 % 10  // 7
const minInt32 int = -0x80000000        // -2147483648
const minLeftDigits int = minInt32 / 10 // -214748364
const minLastDigit int = minInt32 % 10  // -8

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func Reverse(value int) int {
	reversed := 0

	for nextValue := value / 10; nextValue != 0; nextValue /= 10 {
		reversed, value = reversed*10+value%10, nextValue
	}

	if value >= 0 {
		if reversed > maxLeftDigits ||
			reversed == maxLeftDigits && value > maxLastDigit {
			return 0
		}
	} else {
		if reversed < minLeftDigits ||
			reversed == minLeftDigits && value < minLastDigit {
			return 0
		}
	}

	return reversed*10 + value
}
