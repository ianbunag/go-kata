package josephus_permutation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "fmt"

	. "github.com/yvnbunag/go-kata/codewars/46_josephus_permutation"
)

func dotest(items []interface{}, k int, exp []interface{}) {
	Println("input:", items, k)
	var ans = Josephus(items, k)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("46JosephusPermutation", func() {
	It("should handle basic test cases", func() {
		items := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		k := 1
		dotest(items, k, result)
		items = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result = []interface{}{2, 4, 6, 8, 10, 3, 7, 1, 9, 5}
		k = 2
		dotest(items, k, result)
		items = []interface{}{1, 2, 3, 4, 5, 6, 7}
		result = []interface{}{3, 6, 2, 7, 5, 1, 4}
		k = 3
		dotest(items, k, result)
		items = []interface{}{}
		result = []interface{}{}
		k = 3
		dotest(items, k, result)
		items = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		result = []interface{}{10, 7, 8, 13, 5, 4, 12, 11, 3, 15, 14, 9, 1, 6, 2}
		k = 40
		dotest(items, k, result)
	})
})
