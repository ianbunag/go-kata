package rgb_to_hex_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/28_rgb_to_hex"
)

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(RGB(0, 0, 0)).To(Equal("000000"))
		// Expect(RGB(1, 2, 3)).To(Equal("010203"))
		// Expect(RGB(255, 255, 255)).To(Equal("FFFFFF"))
		// Expect(RGB(254, 253, 252)).To(Equal("FEFDFC"))
		// Expect(RGB(-20, 275, 125)).To(Equal("00FF7D"))
	})
})
