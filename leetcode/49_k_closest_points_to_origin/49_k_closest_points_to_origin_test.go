package k_closest_points_to_origin_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/49_k_closest_points_to_origin"
)

var _ = Describe("49KClosestPointsToOrigin", func() {
	It("should get closest points to origin", func() {
		Expect(KClosest([][]int{{1, 3}, {-2, 2}}, 1)).To(ConsistOf([][]int{{-2, 2}}))
		Expect(KClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2)).To(ConsistOf([][]int{{3, 3}, {-2, 4}}))
		Expect(KClosest([][]int{{1, 3}, {-2, 2}, {2, -2}}, 2)).To(ConsistOf([][]int{{-2, 2}, {2, -2}}))
		Expect(KClosest([][]int{{1, 3}, {-2, 2}, {2, -2}}, 3)).To(ConsistOf([][]int{{-2, 2}, {2, -2}, {1, 3}}))
		Expect(KClosest([][]int{{1, 3}, {-2, 2}, {2, -2}}, 4)).To(ConsistOf([][]int{{-2, 2}, {2, -2}, {1, 3}}))
		Expect(KClosest([][]int{{1, 3}, {-2, 2}, {2, -2}}, 0)).To(ConsistOf([][]int{}))
		Expect(KClosest([][]int{{1, 3}, {-2, 2}, {2, -2}}, 5)).To(ConsistOf([][]int{{-2, 2}, {2, -2}, {1, 3}}))
	})
})
