package binomial_coefficient_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test3BinomialCoefficient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "3BinomialCoefficient Suite")
}
