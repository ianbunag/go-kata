package merge_intervals_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/38_merge_intervals"
)

var _ = Describe("38MergeIntervals", func() {
	It("should merge intervals", func() {
		Expect(
			Merge([][]int{{1, 4}, {0, 2}, {3, 5}}),
		).To(Equal([][]int{{0, 5}}))
		Expect(
			Merge([][]int{{1, 4}}),
		).To(Equal([][]int{{1, 4}}))
		Expect(
			Merge([][]int{{1, 4}, {0, 0}}),
		).To(Equal([][]int{{0, 0}, {1, 4}}))
		Expect(
			Merge([][]int{{1, 2}, {3, 4}}),
		).To(Equal([][]int{{1, 2}, {3, 4}}))
		Expect(
			Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}),
		).To(Equal([][]int{{1, 6}, {8, 10}, {15, 18}}))
		Expect(
			Merge([][]int{{1, 4}, {1, 5}}),
		).To(Equal([][]int{{1, 5}}))
	})
})
