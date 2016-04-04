package main_test

import (
	"io/ioutil"
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
	var folder string // Temp folder as input to ../bin/in

	// Clear response
	JustBeforeEach(func() {
		response.Clear()
	})

	BeforeEach(func() {
		err := request.Setup(
			[]byte(`{
				"source": {
					"server": {
						"protocol": "",
						"host": "localhost",
						"port": 1666
					},
					"user": "Joe_Coder",
					"password": "",
					"filespec": {
						"depot": "...",
						"stream": "...",
						"path": "..."
					}
				},
				"version": {"changelist": ""}
			}`))
		Expect(err).To(Not(HaveOccurred()))

		folder, err = ioutil.TempDir("", "")
		Expect(err).To(Not(HaveOccurred()))
		cmd = *exec.Command("../bin/in", folder)
	})

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

	Context("when version is submitted", func() {
		BeforeEach(func() {
			(&request).Version.Changelist = "12099"
			code = 0
		})

		It("should return the fetched version", func() {
			Expect(response.Version.Changelist).To(Equal("12099"))
		})

		It("should sync it in the target folder", func() {

		})
	})
})
