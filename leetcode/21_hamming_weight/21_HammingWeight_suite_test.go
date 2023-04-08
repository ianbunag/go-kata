package hamming_weight_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test21HammingWeight(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "21HammingWeight Suite")
}
