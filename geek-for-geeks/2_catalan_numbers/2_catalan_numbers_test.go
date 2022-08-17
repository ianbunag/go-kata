package catalan_numbers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/geek-for-geeks/2_catalan_numbers"
)

var _ = Describe("2CatalanNumbers", func() {
	It("should get catalan numbers", func() {
		Expect(GetCatalanNumber(0)).To(Equal(1))
		Expect(GetCatalanNumber(1)).To(Equal(1))
		Expect(GetCatalanNumber(2)).To(Equal(2))
		Expect(GetCatalanNumber(3)).To(Equal(5))
		Expect(GetCatalanNumber(4)).To(Equal(14))
		Expect(GetCatalanNumber(5)).To(Equal(42))
		Expect(GetCatalanNumber(6)).To(Equal(132))
		Expect(GetCatalanNumber(7)).To(Equal(429))
		Expect(GetCatalanNumber(8)).To(Equal(1430))
		Expect(GetCatalanNumber(9)).To(Equal(4862))
	})
})
