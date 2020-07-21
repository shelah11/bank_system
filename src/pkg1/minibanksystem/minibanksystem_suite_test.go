package minibanksystem_test

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)
func TestMinibanksystem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Minibanksystem Suite")
}
