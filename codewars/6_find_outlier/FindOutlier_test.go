package find_outlier_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/6_find_outlier"
)

var _ = Describe("FindOutlier", func() {
	It("Example test case", func() {
		Expect(FindOutlier([]int{2, 6, 8, -10, 3})).To(Equal(3))
		Expect(FindOutlier([]int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781})).To(Equal(206847684))
		Expect(FindOutlier([]int{math.MaxInt32, 0, 1})).To(Equal(0))
	})
})
