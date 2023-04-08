package sum_without_plus_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/20_sum_without_plus"
)

var _ = Describe("20SumWithoutPlus", func() {
	It("should add without plus", func() {
		Expect(GetSum(1, 2)).To(Equal(3))
		Expect(GetSum(2, 3)).To(Equal(5))
		Expect(GetSum(20, 30)).To(Equal(50))
		Expect(GetSum(-1, 1)).To(Equal(0))
	})
})
