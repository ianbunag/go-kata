package is_valid_ip_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test12IsValidIp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "12IsValidIp Suite")
}
