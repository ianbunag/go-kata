package common_denominators_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/47_common_denominators"
)

func dotest(a [][]int, exp string) {
	var ans = ConvertFracts(a)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("47CommonDenominators", func() {
	It("Basic tests", func() {
		var lst = [][]int{{3, 6}, {1, 3}, {1, 4}}
		dotest(lst, "(6,12)(4,12)(3,12)")

		lst = [][]int{{1, 2}, {1, 3}, {1, 4}}
		dotest(lst, "(6,12)(4,12)(3,12)")

		lst = [][]int{{69, 130}, {87, 1310}, {30, 40}}
		dotest(lst, "(18078,34060)(2262,34060)(25545,34060)")
	})

	Describe("Fraction", func() {
		It("should simplify", func() {
			fraction := Fraction{}

			fraction = Fraction{30, 40}
			fraction.Simplify()
			Expect(fraction).To(Equal(Fraction{3, 4}))

			fraction = Fraction{69, 130}
			fraction.Simplify()
			Expect(fraction).To(Equal(Fraction{69, 130}))

			fraction = Fraction{1, 2}
			fraction.Simplify()
			Expect(fraction).To(Equal(Fraction{1, 2}))

			fraction = Fraction{2, 4}
			fraction.Simplify()
			Expect(fraction).To(Equal(Fraction{1, 2}))
		})
	})
})
