package hand_of_straights_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/76_hand_of_straights"
)

var _ = Describe("76HandOfStraights", func() {
	It("should return true if the cards can be rearranged", func() {
		Expect(IsNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3)).To(Equal(true))
	})

	It("should return false if the cards cannot be rearrange", func() {
		Expect(IsNStraightHand([]int{1, 2, 3, 4, 5}, 4)).To(Equal(false))
	})
})
