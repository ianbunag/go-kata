package lib_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/yvnbunag/go-kata/lib"
)

var _ = Describe("BinarySearchTree", func() {
	It("should be sorted", func() {
		treeNode := lib.TreeNode{Value: 3}
		treeNode.Insert(2)
		treeNode.Insert(4)
		treeNode.Insert(5)
		treeNode.Insert(1)

		Expect(treeNode.ToList()).To(Equal([]int{1, 2, 3, 4, 5}))
	})

	It("should get node", func() {
		treeNode := lib.TreeNode{Value: 3}
		treeNode.Insert(2)
		treeNode.Insert(4)
		treeNode.Insert(5)
		treeNode.Insert(1)

		Expect(treeNode.Get(1).Value).To(Equal(1))
		Expect(treeNode.Get(2).Value).To(Equal(2))
		Expect(treeNode.Get(3).Value).To(Equal(3))
		Expect(treeNode.Get(4).Value).To(Equal(4))
		Expect(treeNode.Get(5).Value).To(Equal(5))
		Expect(treeNode.Get(6)).To(BeNil())
	})

	It("should traverse next", func() {
		treeNode := lib.TreeNode{Value: 4}
		treeNode.Insert(5)
		treeNode.Insert(3)
		treeNode.Insert(8)

		walkerNode := treeNode.Get(3)
		Expect(walkerNode.Value).To(Equal(3))

		walkerNode = walkerNode.Next()
		Expect(walkerNode.Value).To(Equal(4))

		walkerNode = walkerNode.Next()
		Expect(walkerNode.Value).To(Equal(5))

		walkerNode = walkerNode.Next()
		Expect(walkerNode.Value).To(Equal(8))

		walkerNode = walkerNode.Next()
		Expect(walkerNode).To(BeNil())
	})

	It("should traverse previous", func() {
		treeNode := lib.TreeNode{Value: 4}
		treeNode.Insert(5)
		treeNode.Insert(3)
		treeNode.Insert(8)

		walkerNode := treeNode.Get(8)
		Expect(walkerNode.Value).To(Equal(8))

		walkerNode = walkerNode.Prev()
		Expect(walkerNode.Value).To(Equal(5))

		walkerNode = walkerNode.Prev()
		Expect(walkerNode.Value).To(Equal(4))

		walkerNode = walkerNode.Prev()
		Expect(walkerNode.Value).To(Equal(3))

		walkerNode = walkerNode.Prev()
		Expect(walkerNode).To(BeNil())
	})

	It("should traverse next and previous", func() {
		treeNode := lib.TreeNode{Value: 4}
		treeNode.Insert(5)
		treeNode.Insert(3)
		treeNode.Insert(8)

		walkerNode := treeNode.Get(3)
		Expect(walkerNode.Value).To(Equal(3))

		walkerNode = walkerNode.Next()
		Expect(walkerNode.Value).To(Equal(4))

		walkerNode = walkerNode.Next()
		Expect(walkerNode.Value).To(Equal(5))

		walkerNode = walkerNode.Next()
		Expect(walkerNode.Value).To(Equal(8))

		walkerNode = walkerNode.Prev()
		Expect(walkerNode.Value).To(Equal(5))

		walkerNode = walkerNode.Prev()
		Expect(walkerNode.Value).To(Equal(4))

		walkerNode = walkerNode.Prev()
		Expect(walkerNode.Value).To(Equal(3))
	})
})
