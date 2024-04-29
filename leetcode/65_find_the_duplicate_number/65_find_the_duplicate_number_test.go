package find_the_duplicate_number_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/65_find_the_duplicate_number"
)

var _ = Describe("65FindTheDuplicateNumber", func() {
	It("should find duplicate number", func() {
		Expect(FindDuplicate([]int{1, 3, 4, 2, 2})).To(Equal(2))
		Expect(FindDuplicate([]int{3, 1, 3, 4, 2})).To(Equal(3))
		Expect(FindDuplicate([]int{3, 3, 3, 3, 3})).To(Equal(3))
	})
})
