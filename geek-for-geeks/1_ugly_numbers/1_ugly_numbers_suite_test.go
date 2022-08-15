package ugly_numbers_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test1UglyNumbers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "1UglyNumbers Suite")
}
