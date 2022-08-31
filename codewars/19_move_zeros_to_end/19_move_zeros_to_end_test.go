package move_zeros_to_end_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/19_move_zeros_to_end"
)

var _ = Describe("19MoveZerosToEnd", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1})).To(Equal([]int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0}))
	})
})
