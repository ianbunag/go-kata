package time_based_key_value_store_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test61TimeBasedKeyValueStore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "61TimeBasedKeyValueStore Suite")
}
