package merge_triplets_to_form_target_triplet_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/67_merge_triplets_to_form_target_triplet"
)

var _ = Describe("67MergeTripletsToFormTargetTriplet", func() {
	It("should merge triplets", func() {
		Expect(MergeTriplets([][]int{
			{2, 5, 3},
			{1, 8, 4},
			{1, 7, 5},
		}, []int{2, 7, 5})).To(Equal(true))
		Expect(MergeTriplets([][]int{
			{3, 4, 5},
			{4, 5, 6},
		}, []int{2, 7, 5})).To(Equal(false))
		Expect(MergeTriplets([][]int{
			{2, 5, 3},
			{2, 3, 4},
			{1, 2, 5},
			{5, 2, 3},
		}, []int{5, 5, 5})).To(Equal(true))
		Expect(MergeTriplets([][]int{
			{7, 15, 15},
			{11, 8, 3},
			{5, 3, 4},
			{12, 9, 9},
			{5, 12, 10},
			{7, 15, 10},
			{7, 6, 4},
			{3, 9, 8},
			{2, 13, 1},
			{14, 2, 3},
		}, []int{14, 6, 4})).To(Equal(true))
	})
})
