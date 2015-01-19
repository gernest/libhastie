package libhastie_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLibhastie(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Libhastie Suite")
}
