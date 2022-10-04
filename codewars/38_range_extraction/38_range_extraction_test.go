package range_extraction_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/38_range_extraction"
)

var _ = Describe("38RangeExtraction", func() {
	It("should match the example", func() {
		s := Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20})
		Expect(s).To(Equal("-6,-3-1,3-5,7-11,14,15,17-20"))
	})
})
