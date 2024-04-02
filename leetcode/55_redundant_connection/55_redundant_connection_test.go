package redundant_connection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/55_redundant_connection"
)

var _ = Describe("55RedundantConnection", func() {
	It("should find redundant connection", func() {
		Expect(FindRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}})).To(Equal([]int{2, 3}))
		Expect(FindRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}})).To(Equal([]int{1, 4}))
		Expect(FindRedundantConnection([][]int{{1, 5}, {3, 4}, {3, 5}, {4, 5}, {2, 4}})).To(Equal([]int{4, 5}))
	})

	It("should find redundant connection (v1)", func() {
		Expect(FindRedundantConnectionV1([][]int{{1, 2}, {1, 3}, {2, 3}})).To(Equal([]int{2, 3}))
		Expect(FindRedundantConnectionV1([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}})).To(Equal([]int{1, 4}))
		Expect(FindRedundantConnectionV1([][]int{{1, 5}, {3, 4}, {3, 5}, {4, 5}, {2, 4}})).To(Equal([]int{4, 5}))
	})
})
