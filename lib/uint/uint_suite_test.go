package lib_uint_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUint(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uint Suite")
}
