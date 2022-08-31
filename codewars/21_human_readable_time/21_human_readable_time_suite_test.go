package human_readable_time_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test21HumanReadableTime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "21HumanReadableTime Suite")
}
