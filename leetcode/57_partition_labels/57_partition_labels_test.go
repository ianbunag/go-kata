package partition_labels_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/57_partition_labels"
)

var _ = Describe("57PartitionLabels", func() {
	It("should partition labels", func() {
		Expect(PartitionLabels("a")).To(Equal([]int{1}))
		Expect(PartitionLabels("ab")).To(Equal([]int{1, 1}))
		Expect(PartitionLabels("aba")).To(Equal([]int{3}))
		Expect(PartitionLabels("abab")).To(Equal([]int{4}))
		Expect(PartitionLabels("ababcbacadefegdehijhklij")).To(Equal([]int{9, 7, 8}))
		Expect(PartitionLabels("eccbbbbdec")).To(Equal([]int{10}))
	})
})
