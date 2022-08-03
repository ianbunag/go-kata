package sum_of_intervals_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/17_sum_of_intervals"
)

var _ = Describe("17SumOfIntervals", func() {
	It("Sample tests", func() {
		Expect(SumOfIntervals([][2]int{{1, 5}})).To(Equal(4))
		Expect(SumOfIntervals([][2]int{{1, 1}})).To(Equal(0))
		Expect(SumOfIntervals([][2]int{{1, 5}, {6, 10}})).To(Equal(8))
		Expect(SumOfIntervals([][2]int{{6, 10}, {1, 5}})).To(Equal(8))
		Expect(SumOfIntervals([][2]int{{1, 5}, {1, 5}})).To(Equal(4))
		// Expect(SumOfIntervals([][2]int{{4, 5}, {3, 8}})).To(Equal(5))
		// Expect(SumOfIntervals([][2]int{{4, 5}, {6, 7}, {3, 8}})).To(Equal(5))
		// Expect(SumOfIntervals([][2]int{{2, 4}, {3, 8}, {7, 9}})).To(Equal(7))
		// Expect(SumOfIntervals([][2]int{{1, 4}, {7, 10}, {3, 5}})).To(Equal(7))
	})
})
