package reverse_integer_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/41_reverse_integer"
)

var _ = Describe("41ReverseInteger", func() {
	It("should reverse integer,", func() {
		Expect(Reverse(-2147483648)).To(Equal(0))
		Expect(Reverse(1534236469)).To(Equal(0))
		Expect(Reverse(123)).To(Equal(321))
		Expect(Reverse(-123)).To(Equal(-321))
		Expect(Reverse(120)).To(Equal(21))
	})
})
