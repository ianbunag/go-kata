package square_sum_of_squared_divisors_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/35_square_sum_of_squared_divisors"
)

func dotest(m, n int, exp [][]int) {
	var ans = ListSquared(m, n)
	Expect(ans).To(ConsistOf(exp))
}

var _ = Describe("35SquareSumOfSquaredDivisors", func() {
	It("should handle basic cases", func() {
		dotest(246, 246, [][]int{{246, 84100}})
		dotest(1, 250, [][]int{{1, 1}, {42, 2500}, {246, 84100}})
		dotest(250, 500, [][]int{{287, 84100}})
		dotest(300, 600, [][]int{})
	})
})
