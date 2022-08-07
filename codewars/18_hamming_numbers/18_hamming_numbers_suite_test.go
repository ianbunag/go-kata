package hamming_numbers_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test18HammingNumbers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "18HammingNumbers Suite")
}
