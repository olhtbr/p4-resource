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
					"depot": "...",
					"stream": "...",
					"path": "..."
				}
			},
			"version": {"changelist": ""}
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
			Expect(response[0].Changelist).To(Equal("12104"))
		})
	})

	Context("when version is latest", func() {
		BeforeEach(func() {
			request.Version.Changelist = "12104"
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
				request.Version.Changelist = "12100"
			})

			ValidateList()
		})

		Context("and is not deleted", func() {
			BeforeEach(func() {
				expected = []string{"12103", "12104"}
				request.Version.Changelist = "12099"
			})

			ValidateList()
		})
	})
})
