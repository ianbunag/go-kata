package copy_list_with_random_pointer_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/68_copy_list_with_random_pointer"
)

var _ = Describe("68CopyListWithRandomPointer", func() {
	It("should copy list with random pointer", func() {
		node7 := &Node{Val: 7}
		node13 := &Node{Val: 13}
		node11 := &Node{Val: 11}
		node10 := &Node{Val: 10}
		node1 := &Node{Val: 1}

		node7.Next = node13
		node7.Random = nil
		node13.Next = node11
		node13.Random = node7
		node11.Next = node10
		node11.Random = node1
		node10.Next = node1
		node10.Random = node11
		node1.Next = nil
		node1.Random = node7

		copy := CopyRandomList(node7)

		Expect(copy.Val).To(Equal(7))
		Expect(copy.Random).To(BeNil())
		copy = copy.Next
		Expect(copy.Val).To(Equal(13))
		Expect(copy.Random.Val).To(Equal(7))
		copy = copy.Next
		Expect(copy.Val).To(Equal(11))
		Expect(copy.Random.Val).To(Equal(1))
		copy = copy.Next
		Expect(copy.Val).To(Equal(10))
		Expect(copy.Random.Val).To(Equal(11))
		copy = copy.Next
		Expect(copy.Val).To(Equal(1))
		Expect(copy.Random.Val).To(Equal(7))
		Expect(copy.Next).To(BeNil())
	})

	It("should copy nil list", func() {
		Expect(CopyRandomList(nil)).To(BeNil())
	})
})
