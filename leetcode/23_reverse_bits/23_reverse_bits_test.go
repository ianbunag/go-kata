package reverse_bits_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/leetcode/23_reverse_bits"
)

var _ = Describe("23ReverseBits", func() {
	It("should reverse bits", func() {
		Expect(
			ReverseBits(0b00000010100101000001111010011100),
		).To(Equal(uint32(0b00111001011110000010100101000000)))
		Expect(
			ReverseBits(0b11111111111111111111111111111101),
		).To(Equal(uint32(0b10111111111111111111111111111111)))
	})
})
