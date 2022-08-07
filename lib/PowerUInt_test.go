package lib_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/lib"
)

var _ = Describe("PowerUInt", func() {
	It("should raise to the power of the exponent", func() {
		Expect(PowerUInt(1, 0)).To(Equal(uint(1)))
		Expect(PowerUInt(2, 0)).To(Equal(uint(1)))
		Expect(PowerUInt(1, 1)).To(Equal(uint(1)))
		Expect(PowerUInt(2, 2)).To(Equal(uint(4)))
		Expect(PowerUInt(1, 1)).To(Equal(uint(1)))
		Expect(PowerUInt(2, 3)).To(Equal(uint(8)))
	})
})
