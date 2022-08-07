package hamming_numbers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/18_hamming_numbers"
)

var _ = Describe("18HammingNumbers", func() {
	It("Should handle basic cases", func() {
		Expect(Hammer(1)).To(Equal(uint(1)))
		Expect(Hammer(2)).To(Equal(uint(2)))
		Expect(Hammer(3)).To(Equal(uint(3)))
		Expect(Hammer(4)).To(Equal(uint(4)))
		Expect(Hammer(5)).To(Equal(uint(5)))
		Expect(Hammer(6)).To(Equal(uint(6)))
		Expect(Hammer(7)).To(Equal(uint(8)))
		Expect(Hammer(8)).To(Equal(uint(9)))
		Expect(Hammer(9)).To(Equal(uint(10)))
		Expect(Hammer(10)).To(Equal(uint(12)))
		Expect(Hammer(11)).To(Equal(uint(15)))
		Expect(Hammer(12)).To(Equal(uint(16)))
		Expect(Hammer(13)).To(Equal(uint(18)))
		Expect(Hammer(14)).To(Equal(uint(20)))
		Expect(Hammer(15)).To(Equal(uint(24)))
		Expect(Hammer(16)).To(Equal(uint(25)))
		Expect(Hammer(17)).To(Equal(uint(27)))
		Expect(Hammer(18)).To(Equal(uint(30)))
		Expect(Hammer(500)).To(Equal(uint(937500)))
		Expect(Hammer(600)).To(Equal(uint(2460375)))
		Expect(Hammer(700)).To(Equal(uint(5898240)))
		Expect(Hammer(800)).To(Equal(uint(12754584)))
		Expect(Hammer(13282)).To(Equal(uint(18432000000000000000)))
	})
})
