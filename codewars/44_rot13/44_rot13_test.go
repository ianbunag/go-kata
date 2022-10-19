package rot13_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/44_rot13"
)

var _ = Describe("44Rot13", func() {

	It("should test that the solution returns the correct value", func() {
		Expect(Rot13("a")).To(Equal("n"))
		Expect(Rot13("ab")).To(Equal("no"))
		Expect(Rot13("an")).To(Equal("na"))
		Expect(Rot13("anAN")).To(Equal("naNA"))
		Expect(Rot13("abcdefghijklmnopqrstuvwxyz")).To(Equal("nopqrstuvwxyzabcdefghijklm"))
		Expect(Rot13("ABCDEFGHIJKLMNOPQRSTUVWXYZ")).To(Equal("NOPQRSTUVWXYZABCDEFGHIJKLM"))
		Expect(Rot13("EBG13 rknzcyr.")).To(Equal("ROT13 example."))
		Expect(Rot13("How can you tell an extrovert from an\nintrovert at NSA? Va gur ryringbef,\ngur rkgebireg ybbxf ng gur BGURE thl'f fubrf.")).To(Equal("Ubj pna lbh gryy na rkgebireg sebz na\nvagebireg ng AFN? In the elevators,\nthe extrovert looks at the OTHER guy's shoes."))
		Expect(Rot13("123")).To(Equal("123"))
		Expect(Rot13("Guvf vf npghnyyl gur svefg xngn V rire znqr. Gunaxf sbe svavfuvat vg! :)")).To(Equal("This is actually the first kata I ever made. Thanks for finishing it! :)"))
		Expect(Rot13("@[`{")).To(Equal("@[`{"))
	})
})
