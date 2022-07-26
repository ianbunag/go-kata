package to_camel_case_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test4ToCamelCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "4ToCamelCase Suite")
}
