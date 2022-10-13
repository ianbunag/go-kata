package sum_by_factors_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/41_sum_by_factors"
)

func dotest(lst []int, exp string) {
	var ans = SumOfDivided(lst)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("41SumByFactors", func() {
	It("Basic tests", func() {
		lst1 := []int{12, 15}
		dotest(lst1, "(2 12)(3 27)(5 15)")

		lst2 := []int{15, 21, 24, 30, 45}
		dotest(lst2, "(2 54)(3 135)(5 90)(7 21)")

		dotest([]int{12, -12}, "(2 0)(3 0)")
		dotest([]int{12, -15}, "(2 12)(3 -3)(5 -15)")
		dotest([]int{114, 237, 421}, "(2 114)(3 351)(19 114)(79 237)(421 421)")
		dotest([]int{3366, -4927, -4637, 564, 1242, 1710, -3165, -3867, -1057, 3744, -1359, 774, -1681, 144, 264}, "(2 11808)(3 3417)(5 -1455)(7 -1057)(11 3630)(13 -1183)(17 3366)(19 1710)(23 1242)(41 -1681)(43 774)(47 564)(151 -2416)(211 -3165)(379 -4927)(1289 -3867)(4637 -4637)")
		dotest([]int{-29804, -4209, -28265, -72769, -31744}, "(2 -61548)(3 -4209)(5 -28265)(23 -4209)(31 -31744)(53 -72769)(61 -4209)(1373 -72769)(5653 -28265)(7451 -29804)")
	})
})
