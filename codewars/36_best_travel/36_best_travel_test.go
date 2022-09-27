package best_travel_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/36_best_travel"
)

func dotest(t, k int, ls []int, exp int) {
	var ans = ChooseBestSum(t, k, ls)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("36BestTravel", func() {
	It("should handle basic cases", func() {
		dotest(53, 1, []int{51, 52, 53, 54, 50}, 53)
		var ts = []int{50, 55, 56, 57, 58}
		dotest(163, 3, ts, 163)
		ts = []int{50}
		dotest(163, 3, ts, -1)
		ts = []int{91, 74, 73, 85, 73, 81, 87}
		dotest(230, 3, ts, 228)
		dotest(331, 2, ts, 178)
		dotest(331, 4, ts, 331)
		dotest(331, 5, ts, -1)
	})
})
