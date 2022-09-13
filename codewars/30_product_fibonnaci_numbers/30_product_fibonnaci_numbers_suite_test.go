package product_fibonnaci_numbers_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test30ProductFibonnaciNumbers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "30ProductFibonnaciNumbers Suite")
}
