package main_test

import (
	"encoding/json"
	"os/exec"

	"github.com/olhtbr/p4-resource/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Check executed", func() {
	var jsonBlob []byte
	var cmd *exec.Cmd
	var request models.CheckRequest
	var response models.CheckResponse

	JustBeforeEach(func() {
		err := json.Unmarshal(jsonBlob, &request)
		Expect(err).To(Not(HaveOccurred()))
	})

	BeforeEach(func() {
		jsonBlob = []byte(`{
			"source": {
				"port": {
					"protocol": "ssl",
						"host": "my.perforce.server",
						"port": 1668
				},
				"user": "test.user",
				"ticket": "123456ABCDEF",
				"filespec": {
					"depot": "depot",
					"stream": "stream",
					"path": "..."
				}
			},
			"version": {"changelist": "123456"}
		}`)

		cmd = exec.Command("bin/check")
	})

	Context("when version is omitted", func() {

		BeforeEach(func() {
			request.Version.Changelist = ""
		})

		It("should return the latest version", func() {

			Expect(response).To(HaveLen(1))
			Expect(response[0].Changelist).To(Equal("123456"))
		})
	})
})
