package same_tree_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/leetcode/13_same_tree"
	binary "github.com/yvnbunag/go-kata/lib/binary_tree"
)

var _ = Describe("13SameTree", func() {
	It("should return true for same trees", func() {
		nodeA1 := binary.New(1)
		nodeA2 := binary.New(2)
		nodeA3 := binary.New(3)
		nodeA1.Left = nodeA2
		nodeA1.Right = nodeA3

		nodeB1 := binary.New(1)
		nodeB2 := binary.New(2)
		nodeB3 := binary.New(3)
		nodeB1.Left = nodeB2
		nodeB1.Right = nodeB3

		Expect(IsSameTree(nodeA1, nodeB1)).To(Equal(true))
	})
	It("should return false for different trees", func() {
		nodeA1 := binary.New(1)
		nodeA2 := binary.New(2)
		nodeA1.Left = nodeA2

		nodeB1 := binary.New(1)
		nodeB2 := binary.New(2)
		nodeB1.Right = nodeB2

		Expect(IsSameTree(nodeA1, nodeB1)).To(Equal(false))
	})

	It("should return false for inverse trees", func() {
		nodeA1 := binary.New(1)
		nodeA2 := binary.New(2)
		nodeA3 := binary.New(3)
		nodeA1.Left = nodeA2
		nodeA1.Right = nodeA3

		nodeB1 := binary.New(1)
		nodeB2 := binary.New(2)
		nodeB3 := binary.New(3)
		nodeB1.Left = nodeB3
		nodeB1.Right = nodeB2

		Expect(IsSameTree(nodeA1, nodeB1)).To(Equal(false))
	})
})
