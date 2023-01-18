package best_time_stocks_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/leetcode/6_best_time_stocks"
)

var _ = Describe("6BestTimeStocks", func() {
	It("should give max profit", func() {
		Expect(MaxProfit([]int{7, 1, 5, 3, 6, 4})).To(Equal(5))
		Expect(MaxProfit([]int{7, 6, 4, 3, 1})).To(Equal(0))
		Expect(MaxProfit([]int{2, 1, 4})).To(Equal(3))
		Expect(MaxProfit([]int{2, 4, 1})).To(Equal(2))
	})
})
