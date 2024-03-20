package kth_largest_element_in_a_stream_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/47_kth_largest_element_in_a_stream"
)

var _ = Describe("47KthLargestElementInAStream", func() {
	It("it should return kth largest element in a stream", func() {
		kthLargest := Constructor(3, []int{4, 5, 8, 2})

		Expect(kthLargest.Add(3)).To(Equal(4))
		Expect(kthLargest.Add(5)).To(Equal(5))
		Expect(kthLargest.Add(10)).To(Equal(5))
		Expect(kthLargest.Add(9)).To(Equal(8))
		Expect(kthLargest.Add(4)).To(Equal(8))
	})
})
