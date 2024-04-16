package time_based_key_value_store_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/61_time_based_key_value_store"
)

var _ = Describe("61TimeBasedKeyValueStore", func() {
	It("should set and get value", func() {
		obj := Constructor()
		obj.Set("foo", "bar", 1)
		Expect(obj.Get("foo", 1)).To(Equal("bar"))
		Expect(obj.Get("foo", 3)).To(Equal("bar"))
		obj.Set("foo", "bar2", 4)
		Expect(obj.Get("foo", 4)).To(Equal("bar2"))
		Expect(obj.Get("foo", 5)).To(Equal("bar2"))
	})

	It("should return empty string if key is not found", func() {
		obj := Constructor()
		Expect(obj.Get("foo", 1)).To(Equal(""))
	})

	It("should return empty string if a lower timestamp is not found", func() {
		obj := Constructor()
		obj.Set("foo", "bar", 5)
		Expect(obj.Get("foo", 4)).To(Equal(""))
	})
})
