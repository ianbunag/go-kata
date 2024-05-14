package daily_temperatures_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test70DailyTemperatures(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "70DailyTemperatures Suite")
}
