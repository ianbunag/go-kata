package sum_digit_power_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/9_sum_digit_power"
)

func dotest(a, b uint64, expected []uint64) {
	actual := SumDigPow(a, b)
	if len(expected) == 0 {
		Expect(actual).To(BeEmpty(), "With a = %d, b = %d", a, b)
	} else {
		Expect(actual).To(Equal(expected), "With a = %d, b = %d", a, b)
	}
}

var _ = Describe("9SumDigitPower", func() {
	It("Sample tests", func() {
		dotest(90, 100, nil)
		dotest(1, 10, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9})
		dotest(1, 100, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 89})
		dotest(10, 89, []uint64{89})
		dotest(10, 100, []uint64{89})
		dotest(89, 135, []uint64{89, 135})
		dotest(12157692622039623433, 12157692622039625610, []uint64{12157692622039623539})
	})
})
