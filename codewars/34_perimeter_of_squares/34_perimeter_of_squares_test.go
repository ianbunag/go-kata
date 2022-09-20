package perimeter_of_squares_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/34_perimeter_of_squares"
)

func dotest(n int, exp int) {
	var ans = Perimeter(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("34PerimeterOfSquares", func() {
	It("should handle basic cases", func() {
		dotest(5, 80)
		dotest(7, 216)
		dotest(20, 114624)
		dotest(30, 14098308)
	})
})
