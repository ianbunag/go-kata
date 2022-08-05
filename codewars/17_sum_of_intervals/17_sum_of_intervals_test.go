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
		Expect(SumOfIntervals([][2]int{{4, 5}, {3, 8}})).To(Equal(5))
		Expect(SumOfIntervals([][2]int{{4, 5}, {6, 7}, {3, 8}})).To(Equal(5))
		Expect(SumOfIntervals([][2]int{{2, 4}, {3, 8}, {7, 9}})).To(Equal(7))
		Expect(SumOfIntervals([][2]int{{3, 5}, {7, 8}, {10, 11}, {13, 15}, {4, 14}})).To(Equal(12))
		Expect(SumOfIntervals([][2]int{{1, 4}, {7, 10}, {3, 5}})).To(Equal(7))
		Expect(SumOfIntervals([][2]int{{1, 4}, {6, 7}, {2, 9}})).To(Equal(8))
		Expect(SumOfIntervals([][2]int{{6, 7}, {2, 9}, {1, 4}})).To(Equal(8))
		Expect(SumOfIntervals([][2]int{{2, 9}, {1, 4}, {6, 7}})).To(Equal(8))
		Expect(SumOfIntervals([][2]int{{-3, -1}})).To(Equal(2))
		Expect(SumOfIntervals([][2]int{{-3, 1}})).To(Equal(4))
		Expect(SumOfIntervals([][2]int{{42, 69}, {41, 69}})).To(Equal(28))
		Expect(SumOfIntervals([][2]int{
			{42, 69}, {41, 69}, {42, 51}, {-34, 50}, {-51, -18}, {46, 93}, {-25, 78},
			{9, 93}, {-21, 27}, {50, 69}, {-94, 54}, {-39, 39}, {4, 86}, {33, 71},
			{-100, -31}, {13, 42}, {85, 89}, {-47, 24}, {23, 61}, {-98, 18}, {-78, 9},
			{-18, 33},
		})).To(Equal(193))
	})
})
