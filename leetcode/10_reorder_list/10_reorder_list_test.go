package reorder_list_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/10_reorder_list"
	. "github.com/yvnbunag/go-kata/lib/list_node"
)

var _ = Describe("10ReorderList", func() {
	It("should reorder list", func() {
		Expect(ReorderList(nil)).To(BeNil())
		Expect(ReorderList(NewNodeFromList([]int{1, 2, 3, 4, 5})).ToList()).To(Equal([]int{1, 5, 2, 4, 3}))
		Expect(ReorderList(NewNodeFromList([]int{1, 2, 3, 4})).ToList()).To(Equal([]int{1, 4, 2, 3}))
	})
})
