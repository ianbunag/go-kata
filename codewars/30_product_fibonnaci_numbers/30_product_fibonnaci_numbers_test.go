package product_fibonnaci_numbers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/30_product_fibonnaci_numbers"
)

func dotest(prod uint64, exp [3]uint64) {
	var ans = ProductFib(prod)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("30ProductFibonnaciNumbers", func() {
	It("should handle basic cases", func() {
		dotest(4895, [3]uint64{55, 89, 1})
		dotest(5895, [3]uint64{89, 144, 0})
		dotest(74049690, [3]uint64{6765, 10946, 1})
		dotest(84049690, [3]uint64{10946, 17711, 0})
	})
})
