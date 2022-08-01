package sum_from_list_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/14_sum_from_list"
)

var _ = Describe("14SumFromList", func() {
	It("Basic tests", func() {
		Expect(TwoSum([]int{1, 2, 3}, 4)).To(Equal([2]int{0, 2}))
		Expect(TwoSum([]int{1234, 5678, 9012}, 14690)).To(Equal([2]int{1, 2}))
		Expect(TwoSum([]int{2, 2, 3}, 4)).To(Equal([2]int{0, 1}))
	})
})
