package morse_code_advanced_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/40_morse_code_advanced"
)

var _ = Describe("40MorseCodeAdvanced", func() {
	It("Example from description", func() {
		Expect(DecodeMorse(DecodeBits("11111100111111"))).To(Equal("M"))
		Expect(DecodeMorse(DecodeBits("10001"))).To(Equal("EE"))
		Expect(DecodeMorse(DecodeBits("1"))).To(Equal("E"))
		Expect(DecodeMorse(DecodeBits("111"))).To(Equal("E"))
		Expect(DecodeMorse(DecodeBits("1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"))).To(Equal("HEY JUDE"))
		Expect(DecodeMorse(DecodeBits("01110"))).To(Equal("E"))
	})
})
