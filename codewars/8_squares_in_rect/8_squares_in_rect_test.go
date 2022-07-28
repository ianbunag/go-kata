package squares_in_rect_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/8_squares_in_rect"
)

func dotest(lng int, wdth int, exp []int) {
	var ans = SquaresInRect(lng, wdth)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("8SquaresInRect", func() {
	It("should handle basic cases", func() {
		dotest(5, 5, nil)
		dotest(5, 3, []int{3, 2, 1, 1})
		dotest(3, 5, []int{3, 2, 1, 1})
		dotest(20, 14, []int{14, 6, 6, 2, 2, 2})
	})
})
