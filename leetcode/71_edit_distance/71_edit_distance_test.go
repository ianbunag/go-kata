package edit_distance_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/71_edit_distance"
)

var _ = Describe("71EditDistance", func() {
	It("should return minimum number of operations required", func() {
		Expect(MinDistance("ba", "bars")).To(Equal(2))
		Expect(MinDistance("bars", "ba")).To(Equal(2))
		Expect(MinDistance("bar", "bar")).To(Equal(0))
		Expect(MinDistance("horse", "ros")).To(Equal(3))
		Expect(MinDistance("intention", "execution")).To(Equal(5))
		Expect(MinDistance("plasma", "altruism")).To(Equal(6))
	})
})
