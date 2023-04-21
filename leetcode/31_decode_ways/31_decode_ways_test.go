package decode_ways_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/31_decode_ways"
)

var _ = Describe("31DecodeWays", func() {
	It("should decode ways", func() {
		Expect(NumDecodings("11106")).To(Equal(2))
		Expect(NumDecodings("226")).To(Equal(3))
		Expect(NumDecodings("06")).To(Equal(0))
		Expect(NumDecodings("12")).To(Equal(2))
		Expect(NumDecodings("111111111111111111111111111111111111111111111")).To(Equal(1836311903))
	})
})
