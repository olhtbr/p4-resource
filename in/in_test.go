package main_test

import (
	"os/exec"

	"github.com/olhtbr/p4-resource/models"
	"github.com/olhtbr/p4-resource/shared"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("In executed", func() {
	var cmd exec.Cmd
	var code int
	var request models.InRequest
	var response models.InResponse

	shared.Setup(&cmd, &request, "../bin/in")
	shared.Run(&cmd, &request, &response, &code)

	Context("when version is omitted", func() {
		BeforeEach(func() {
			code = 1
		})

		It("should exit with error", func() {
			Expect(response.Version.Changelist).To(BeEmpty())
		})
	})

	Context("when version does not exist", func() {
		BeforeEach(func() {
			(&request).Version.Changelist = "12500"
			code = 1
		})

		It("should exit with error", func() {
			Expect(response.Version.Changelist).To(BeEmpty())
		})
	})

	Context("when version is pending", func() {
		BeforeEach(func() {
			(&request).Version.Changelist = "12100"
			code = 1
		})

		It("should exit with error", func() {
			Expect(response.Version.Changelist).To(BeEmpty())
		})
	})
})
