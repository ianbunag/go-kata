package duplicate_count_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/7_duplicate_count"
)

func dotest(prod string, exp int) {
	var ans = CountDuplicate(prod)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("CountDuplicate", func() {
	It("should handle basic cases", func() {
		dotest("abcde", 0)
		dotest("abcdea", 1)
		dotest("abcdeaB11", 3)
		dotest("indivisibility", 1)
	})
})
