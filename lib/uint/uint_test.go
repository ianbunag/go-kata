package lib_uint_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	lib_uint "github.com/yvnbunag/go-kata/lib/uint"
)

var _ = Describe("Uint", func() {
	Context("Power", func() {
		It("should raise to the power of the exponent", func() {
			Expect(lib_uint.Power(1, 0)).To(Equal(uint(1)))
			Expect(lib_uint.Power(2, 0)).To(Equal(uint(1)))
			Expect(lib_uint.Power(1, 1)).To(Equal(uint(1)))
			Expect(lib_uint.Power(2, 2)).To(Equal(uint(4)))
			Expect(lib_uint.Power(1, 1)).To(Equal(uint(1)))
			Expect(lib_uint.Power(2, 3)).To(Equal(uint(8)))
		})
	})

	Context("Min", func() {
		It("should return min", func() {
			Expect(lib_uint.Min(1, 0)).To(Equal(uint(0)))
			Expect(lib_uint.Min(2, 1)).To(Equal(uint(1)))
			Expect(lib_uint.Min(2, 2)).To(Equal(uint(2)))
		})
	})
})
