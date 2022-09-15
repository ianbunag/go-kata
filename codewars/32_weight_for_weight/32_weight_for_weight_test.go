package weight_for_weight_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/32_weight_for_weight"
)

func dotest(s string, exp string) {
	var ans = OrderWeight(s)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("32WeightForWeight", func() {
	It("should handle basic cases", func() {
		dotest("99 100", "100 99")
		dotest("90 180", "180 90")
		dotest("", "")
		dotest("103 123 4444 99 2000", "2000 103 123 4444 99")
		dotest("2000 10003 1234000 44444444 9999 11 11 22 123", "11 11 2000 10003 22 123 1234000 44444444 9999")
	})
})
