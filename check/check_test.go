package main_test

import (
	"os/exec"

	"github.com/olhtbr/p4-resource/models"
	"github.com/olhtbr/p4-resource/shared"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Check executed", func() {
	var cmd exec.Cmd
	var code int
	var request models.CheckRequest
	var response models.CheckResponse

	shared.Setup(&cmd, &request, "../bin/check")
	shared.Run(&cmd, &request, &response, &code)

	Context("when version is omitted", func() {
		BeforeEach(func() {
			(&request).Version.Changelist = ""
		})

		It("should return the latest version", func() {
			Expect(response).To(HaveLen(1))
			Expect(response[0].Changelist).To(Equal("12104"))
		})
	})

	Context("when version is latest", func() {
		BeforeEach(func() {
			(&request).Version.Changelist = "12104"
		})

		It("should return an empty list", func() {
			Expect(response).To(HaveLen(0))
		})
	})

	Context("when version is not the latest", func() {
		var expected []string

		ValidateList := func() {
			It("should return a list of versions", func() {
				Expect(response).To(HaveLen(2))
				Expect(response[0].Changelist).To(Equal(expected[0]))
				Expect(response[1].Changelist).To(Equal(expected[1]))
			})
		}

		Context("and is deleted", func() {
			BeforeEach(func() {
				expected = []string{"12103", "12104"}
				(&request).Version.Changelist = "12100"
			})

			ValidateList()
		})

		Context("and is not deleted", func() {
			BeforeEach(func() {
				expected = []string{"12103", "12104"}
				(&request).Version.Changelist = "12099"
			})

			ValidateList()
		})
	})
})
