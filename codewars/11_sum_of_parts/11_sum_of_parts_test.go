package sum_of_parts_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/11_sum_of_parts"
)

var _ = Describe("11SumOfParts", func() {
	It("Sample tests", func() {
		Expect(PartsSums([]uint64{})).To(Equal([]uint64{0}))
		Expect(PartsSums([]uint64{0, 1, 3, 6, 10})).To(Equal([]uint64{20, 20, 19, 16, 10, 0}))
		Expect(PartsSums([]uint64{1, 2, 3, 4, 5, 6})).To(Equal([]uint64{21, 20, 18, 15, 11, 6, 0}))
		Expect(PartsSums([]uint64{744125, 935, 407, 454, 430, 90, 144, 6710213, 889, 810, 2579358})).To(Equal([]uint64{10037855, 9293730, 9292795, 9292388, 9291934, 9291504, 9291414, 9291270, 2581057, 2580168, 2579358, 00}))
	})
})
