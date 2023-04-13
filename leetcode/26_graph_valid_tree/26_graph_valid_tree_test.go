package graph_valid_tree_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/26_graph_valid_tree"
)

var _ = Describe("26GraphValidTree", func() {
	It("should check if tree is valid", func() {
		validEdges := [][]int{
			{0, 1},
			{0, 2},
			{0, 3},
			{1, 4},
		}
		Expect(ValidTree(5, validEdges)).To(Equal(true))

		outOfOrderEdges := [][]int{
			{1, 0},
		}
		Expect(ValidTree(2, outOfOrderEdges)).To(Equal(true))

		circularEdges := [][]int{
			{0, 1},
			{1, 2},
			{2, 3},
			{1, 3},
			{1, 4},
		}
		Expect(ValidTree(5, circularEdges)).To(Equal(false))

		disconnectedEdges := [][]int{{0, 1}}
		Expect(ValidTree(3, disconnectedEdges)).To(Equal(false))
	})
})
