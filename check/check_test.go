package main_test

import (
	"github.com/olhtbr/p4-resource/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Check executed", func() {
	var response models.CheckResponse

	Context("when version is omitted", func() {
		It("should return the latest version", func() {
			Expect(response).To(HaveLen(1))
			Expect(response[0].Changelist).To(Equal("123456"))
		})
	})
})
