package count_ip_addresses_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/43_count_ip_addresses"
)

var _ = Describe("43CountIpAddresses", func() {
	It("should handle basic cases", func() {
		Expect(IpsBetween("10.0.0.0", "10.0.0.50")).To(Equal(50))
		Expect(IpsBetween("20.0.0.10", "20.0.1.0")).To(Equal(246))
	})
})
