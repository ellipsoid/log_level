package log_level_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLevelLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Level Log Suite")
}
