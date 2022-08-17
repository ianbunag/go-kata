package catalan_numbers_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test2CatalanNumbers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "2CatalanNumbers Suite")
}
