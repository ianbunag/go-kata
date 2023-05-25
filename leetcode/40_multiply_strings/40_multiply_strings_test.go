package multiply_strings_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/40_multiply_strings"
)

var _ = Describe("40MultiplyStrings", func() {
	Describe("Multiply", func() {
		It("should multiply", func() {
			Expect(Multiply("2", "3")).To(Equal("6"))
			Expect(Multiply("10", "1")).To(Equal("10"))
			Expect(Multiply("10", "2")).To(Equal("20"))
			Expect(Multiply("20", "2")).To(Equal("40"))
			Expect(Multiply("90", "2")).To(Equal("180"))
			Expect(Multiply("99", "99")).To(Equal("9801"))
			Expect(Multiply("123", "456")).To(Equal("56088"))
			Expect(Multiply("9133", "0")).To(Equal("0"))
		})
	})

	Describe("Operatable", func() {
		It("should add", func() {
			Expect(NewOperatable("1").Add(NewOperatable("1")).String()).To(Equal("2"))
			Expect(NewOperatable("4").Add(NewOperatable("5")).String()).To(Equal("9"))
			Expect(NewOperatable("5").Add(NewOperatable("5")).String()).To(Equal("10"))
			Expect(NewOperatable("9").Add(NewOperatable("9")).String()).To(Equal("18"))
			Expect(NewOperatable("10").Add(NewOperatable("10")).String()).To(Equal("20"))
			Expect(NewOperatable("10").Add(NewOperatable("1")).String()).To(Equal("11"))
			Expect(NewOperatable("1").Add(NewOperatable("10")).String()).To(Equal("11"))
			Expect(NewOperatable("99").Add(NewOperatable("99")).String()).To(Equal("198"))
			Expect(NewOperatable("99").Add(NewOperatable("99")).String()).To(Equal("198"))
			Expect(NewOperatable("555").Add(NewOperatable("777")).String()).To(Equal("1332"))
		})

		It("should add without mutating instances", func() {
			one, two := NewOperatable("1"), NewOperatable("2")

			Expect(one.Add(two).String()).To(Equal("3"))
			Expect(one.String()).To(Equal("1"))
			Expect(two.String()).To(Equal("2"))
		})

		It("should multiply", func() {
			Expect(NewOperatable("99").Multiply(NewOperatable("9")).String()).To(Equal("891"))
			Expect(NewOperatable("99").Multiply(NewOperatable("99")).String()).To(Equal("9801"))
			Expect(NewOperatable("10").Multiply(NewOperatable("1")).String()).To(Equal("10"))
			Expect(NewOperatable("2").Multiply(NewOperatable("3")).String()).To(Equal("6"))
			Expect(NewOperatable("10").Multiply(NewOperatable("2")).String()).To(Equal("20"))
			Expect(NewOperatable("20").Multiply(NewOperatable("2")).String()).To(Equal("40"))
			Expect(NewOperatable("90").Multiply(NewOperatable("2")).String()).To(Equal("180"))
			Expect(NewOperatable("2").Multiply(NewOperatable("90")).String()).To(Equal("180"))
			Expect(NewOperatable("123").Multiply(NewOperatable("456")).String()).To(Equal("56088"))
		})

		It("should multiply without mutating instances", func() {
			two, three := NewOperatable("2"), NewOperatable("3")

			Expect(two.Multiply(three).String()).To(Equal("6"))
			Expect(two.String()).To(Equal("2"))
			Expect(three.String()).To(Equal("3"))
		})

		It("should shift left", func() {
			Expect(NewOperatable("1").ShiftLeft(1).String()).To(Equal("10"))
			Expect(NewOperatable("1").ShiftLeft(2).String()).To(Equal("100"))
			Expect(NewOperatable("11").ShiftLeft(3).String()).To(Equal("11000"))
		})
	})

	It("should multiply digits", func() {
		Expect(MultiplyDigits("0", "0"), "0")
		Expect(MultiplyDigits("0", "1"), "0")
		Expect(MultiplyDigits("1", "1"), "1")
		Expect(MultiplyDigits("1", "2"), "2")
		Expect(MultiplyDigits("5", "5"), "25")
		Expect(MultiplyDigits("9", "9"), "81")
	})

	It("should add digits", func() {
		Expect(AddDigits("0", "0")).To(Equal("0"))
		Expect(AddDigits("0", "1")).To(Equal("1"))
		Expect(AddDigits("1", "1")).To(Equal("2"))
		Expect(AddDigits("1", "2")).To(Equal("3"))
		Expect(AddDigits("5", "5")).To(Equal("10"))
		Expect(AddDigits("9", "9")).To(Equal("18"))
	})
})
