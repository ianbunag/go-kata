package counting_bits_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/22_counting_bits"
)

var _ = Describe("22CountingBits", func() {
	It("should count bits", func() {
		Expect(CountBits(2)).To(Equal([]int{0, 1, 1}))
		Expect(CountBits(5)).To(Equal([]int{0, 1, 1, 2, 1, 2}))
	})
})
