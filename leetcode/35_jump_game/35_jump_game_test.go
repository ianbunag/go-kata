package jump_game_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/35_jump_game"
)

var _ = Describe("35JumpGame", func() {
	It("should return true if last index is reachable", func() {
		Expect(CanJump([]int{2, 3, 0, 1, 4})).To(Equal(true))
		Expect(CanJump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0})).To(Equal(true))
		Expect(CanJump([]int{1, 1, 2, 2, 0, 1, 1})).To(Equal(true))
		Expect(CanJump([]int{3, 2, 1, 0, 4})).To(Equal(false))
		Expect(CanJump([]int{4, 3, 2, 1})).To(Equal(true))
		Expect(CanJump([]int{2, 5, 0, 0})).To(Equal(true))
		Expect(CanJump([]int{3, 2, 1})).To(Equal(true))
		Expect(CanJump([]int{2, 0, 0})).To(Equal(true))
		Expect(CanJump([]int{2, 0})).To(Equal(true))
		Expect(CanJump([]int{2, 3, 1, 1, 4})).To(Equal(true))
		Expect(CanJump([]int{0})).To(Equal(true))
		Expect(CanJump([]int{1})).To(Equal(true))
	})
})
