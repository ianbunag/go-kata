package multiples_of_3_or_5_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/22_multiples_of_3_or_5"
)

var _ = Describe("22MultiplesOf3Or5", func() {
	It("should handle basic cases", func() {
		Expect(Multiple3And5(2)).To(Equal(0))
		Expect(Multiple3And5(-1)).To(Equal(0))
		Expect(Multiple3And5(0)).To(Equal(0))
		Expect(Multiple3And5(10)).To(Equal(23))
		Expect(Multiple3And5(20)).To(Equal(78))
		Expect(Multiple3And5(200)).To(Equal(9168))
	})
})
