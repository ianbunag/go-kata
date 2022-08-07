package lib_slice_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	lib_slice "github.com/yvnbunag/go-kata/lib/slice"
)

var _ = Describe("Slice", func() {
	Context("Reverse", func() {
		It("should reverse slice", func() {
			Expect(lib_slice.Reverse([]int{1, 2, 3})).To(Equal([]int{3, 2, 1}))
		})
	})
})
