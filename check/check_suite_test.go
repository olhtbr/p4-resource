package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var binary string

var _ = BeforeSuite(func() {
	var err error

	binary, err = gexec.Build("github.com/olhtbr/p4-resource/check")
	Expect(err).To(Not(HaveOccurred()))
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func TestCheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Check Suite")
}
