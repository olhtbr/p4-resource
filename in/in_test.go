package main_test

import (
	"encoding/json"
	"os/exec"

	"github.com/olhtbr/p4-resource/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("In executed", func() {
	var cmd *exec.Cmd
	var request models.InRequest
	var response models.InResponse
	var code int

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
					"depot": "...",
					"stream": "...",
					"path": "..."
				}
			},
			"version": {"changelist": ""}
		}`)

		err := json.Unmarshal(jsonBlob, &request)
		Expect(err).To(Not(HaveOccurred()))

		cmd = exec.Command("../bin/in")
	})

	JustBeforeEach(func() {
		stdin, err := cmd.StdinPipe()
		Expect(err).To(Not(HaveOccurred()))

		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).To(Not(HaveOccurred()))

		err = json.NewEncoder(stdin).Encode(request)
		Expect(err).To(Not(HaveOccurred()))

		Eventually(session).Should(gexec.Exit(code))

		if code == 0 {
			err = json.Unmarshal(session.Out.Contents(), &response)
			Expect(err).To(Not(HaveOccurred()))
		}
	})

	Context("when version is omitted", func() {
		BeforeEach(func() {
			code = 1
		})

		It("should exit with error", func() {
			Expect(response.Version.Changelist).To(BeEmpty())
		})
	})
})
