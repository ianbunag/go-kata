package missing_number_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/leetcode/24_missing_number"
)

var _ = Describe("24MissingNumber", func() {
	It("should get missing number", func() {
		Expect(MissingNumber([]int{3, 0, 1})).To(Equal(2))
		Expect(MissingNumber([]int{0, 1})).To(Equal(2))
		Expect(MissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})).To(Equal(8))
	})
})
