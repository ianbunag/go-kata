package hamming_weight_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/leetcode/21_hamming_weight"
)

var _ = Describe("21HammingWeight", func() {
	It("should get hamming weight", func() {
		Expect(HammingWeight(00000000000000000000000000001011)).To(Equal(3))
		Expect(HammingWeight(00000000000000000000000010000000)).To(Equal(1))
	})
})
