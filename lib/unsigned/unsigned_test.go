package lib_unsigned_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	lib_unsigned "github.com/yvnbunag/go-kata/lib/unsigned"
)

var _ = Describe("Uint", func() {
	Context("Power", func() {
		It("should raise to the power of the exponent", func() {
			Expect(lib_unsigned.Power[uint](1, 0)).To(Equal(uint(1)))
			Expect(lib_unsigned.Power[uint](2, 0)).To(Equal(uint(1)))
			Expect(lib_unsigned.Power[uint](1, 1)).To(Equal(uint(1)))
			Expect(lib_unsigned.Power[uint](2, 2)).To(Equal(uint(4)))
			Expect(lib_unsigned.Power[uint](1, 1)).To(Equal(uint(1)))
			Expect(lib_unsigned.Power[uint](2, 3)).To(Equal(uint(8)))
		})
	})

	Context("Min", func() {
		It("should return min", func() {
			Expect(lib_unsigned.Min[uint](1, 0)).To(Equal(uint(0)))
			Expect(lib_unsigned.Min[uint](2, 1)).To(Equal(uint(1)))
			Expect(lib_unsigned.Min[uint](2, 2)).To(Equal(uint(2)))
		})
	})
})
