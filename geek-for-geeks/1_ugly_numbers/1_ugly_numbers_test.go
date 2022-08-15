package ugly_numbers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/geek-for-geeks/1_ugly_numbers"
)

var _ = Describe("1UglyNumbers", func() {
	It("should find ugly number", func() {
		Expect(GetUglyNumber(1)).To(Equal(1))
		Expect(GetUglyNumber(2)).To(Equal(2))
		Expect(GetUglyNumber(3)).To(Equal(3))
		Expect(GetUglyNumber(4)).To(Equal(4))
		Expect(GetUglyNumber(5)).To(Equal(5))
		Expect(GetUglyNumber(6)).To(Equal(6))
		Expect(GetUglyNumber(7)).To(Equal(8))
		Expect(GetUglyNumber(8)).To(Equal(9))
		Expect(GetUglyNumber(9)).To(Equal(10))
		Expect(GetUglyNumber(10)).To(Equal(12))
		Expect(GetUglyNumber(11)).To(Equal(15))
	})
})
