package get_sum_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/3_get_sum"
)

var _ = Describe("GetSum", func() {
	It("Sample tests", func() {
		Expect(GetSum(-5, -5)).To(Equal(-5))
		Expect(GetSum(0, 0)).To(Equal(0))
		Expect(GetSum(-17, -17)).To(Equal(-17))
		Expect(GetSum(17, 17)).To(Equal(17))
		Expect(GetSum(0, 1)).To(Equal(1))
		Expect(GetSum(1, 2)).To(Equal(3))
		Expect(GetSum(5, -1)).To(Equal(14))
		Expect(GetSum(505, 4)).To(Equal(127759))
		Expect(GetSum(321, 123)).To(Equal(44178))
		Expect(GetSum(0, -1)).To(Equal(-1))
		Expect(GetSum(-50, 0)).To(Equal(-1275))
		Expect(GetSum(-1, -5)).To(Equal(-15))
		Expect(GetSum(-505, 4)).To(Equal(-127755))
		Expect(GetSum(-321, 123)).To(Equal(-44055))
		Expect(GetSum(-5, -1)).To(Equal(-15))
		Expect(GetSum(5, 1)).To(Equal(15))
	})
})
