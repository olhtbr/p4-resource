package main_test

import (
	"encoding/json"
	"os/exec"

	"github.com/olhtbr/p4-resource/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Check executed", func() {
	var cmd *exec.Cmd
	var request models.CheckRequest
	var response models.CheckResponse

	BeforeEach(func() {
		jsonBlob := []byte(`{
			"source": {
				"server": {
					"protocol": "",
					"host": "localhost",
					"port": 1666
				},
				"user": "Joe_Coder",
				"password": "",
				"filespec": {
					"depot": "HR",
					"stream": "draft",
					"path": ""
				}
			},
			"version": {"changelist": "123456"}
		}`)

		err := json.Unmarshal(jsonBlob, &request)
		Expect(err).To(Not(HaveOccurred()))

		cmd = exec.Command("../bin/check")
	})

	JustBeforeEach(func() {
		stdin, err := cmd.StdinPipe()
		Expect(err).To(Not(HaveOccurred()))

		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).To(Not(HaveOccurred()))

		err = json.NewEncoder(stdin).Encode(request)
		Expect(err).To(Not(HaveOccurred()))

		Eventually(session).Should(gexec.Exit(0))

		err = json.Unmarshal(session.Out.Contents(), &response)
		Expect(err).To(Not(HaveOccurred()))
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
