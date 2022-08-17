package binomial_coefficient_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/geek-for-geeks/3_binomial_coefficient"
)

var _ = Describe("3BinomialCoefficient", func() {
	It("should get binomial coefficient", func() {
		Expect(GetBinomialCoefficient(4, 2)).To(Equal(6))
		Expect(GetBinomialCoefficient(5, 2)).To(Equal(10))

		Expect(GetBinomialCoefficient(0, 0)).To(Equal(1))
		Expect(GetBinomialCoefficient(0, 1)).To(Equal(0))
		Expect(GetBinomialCoefficient(0, 2)).To(Equal(0))
		Expect(GetBinomialCoefficient(0, 3)).To(Equal(0))
		Expect(GetBinomialCoefficient(0, 4)).To(Equal(0))
		Expect(GetBinomialCoefficient(0, 5)).To(Equal(0))

		Expect(GetBinomialCoefficient(1, 0)).To(Equal(1))
		Expect(GetBinomialCoefficient(1, 1)).To(Equal(1))
		Expect(GetBinomialCoefficient(1, 2)).To(Equal(0))
		Expect(GetBinomialCoefficient(1, 3)).To(Equal(0))
		Expect(GetBinomialCoefficient(1, 4)).To(Equal(0))
		Expect(GetBinomialCoefficient(1, 5)).To(Equal(0))

		Expect(GetBinomialCoefficient(2, 0)).To(Equal(1))
		Expect(GetBinomialCoefficient(2, 1)).To(Equal(2))
		Expect(GetBinomialCoefficient(2, 2)).To(Equal(1))
		Expect(GetBinomialCoefficient(2, 3)).To(Equal(0))
		Expect(GetBinomialCoefficient(2, 4)).To(Equal(0))
		Expect(GetBinomialCoefficient(2, 5)).To(Equal(0))

		Expect(GetBinomialCoefficient(3, 0)).To(Equal(1))
		Expect(GetBinomialCoefficient(3, 1)).To(Equal(3))
		Expect(GetBinomialCoefficient(3, 2)).To(Equal(3))
		Expect(GetBinomialCoefficient(3, 3)).To(Equal(1))
		Expect(GetBinomialCoefficient(3, 4)).To(Equal(0))
		Expect(GetBinomialCoefficient(3, 5)).To(Equal(0))

		Expect(GetBinomialCoefficient(4, 0)).To(Equal(1))
		Expect(GetBinomialCoefficient(4, 1)).To(Equal(4))
		Expect(GetBinomialCoefficient(4, 2)).To(Equal(6))
		Expect(GetBinomialCoefficient(4, 3)).To(Equal(4))
		Expect(GetBinomialCoefficient(4, 4)).To(Equal(1))
		Expect(GetBinomialCoefficient(4, 5)).To(Equal(0))
	})
})
