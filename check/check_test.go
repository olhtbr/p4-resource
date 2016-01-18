package main_test

import (
	"encoding/json"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/olhtbr/p4-resource/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Check executed", func() {
	var cmd *exec.Cmd
	var request models.CheckRequest
	var response models.CheckResponse
	var cl string

	GetChangelistCounter := func() uint64 {
		cmd := exec.Command("p4", "-z", "tag", "-u", "Joe_Coder", "-p", "localhost:1666", "counter", "change")
		out, err := cmd.Output()
		Expect(err).To(Not(HaveOccurred()))
		Expect(out).To(Not(BeNil()))
		Expect(out).To(Not(BeEmpty()))

		re := regexp.MustCompile(`\.\.\.\s+value\s+(\d+)`)
		match := re.FindStringSubmatch(string(out))
		Expect(len(match)).To(Equal(2))

		counter, err := strconv.ParseUint(match[1], 10, 16)
		Expect(err).To(Not(HaveOccurred()))

		return counter
	}

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
					"path": "..."
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
			request.Source.Filespec.Depot = "..."
			request.Source.Filespec.Stream = "..."
			request.Source.Filespec.Path = "..."

			// Counter is latest changelist + 1
			counter := GetChangelistCounter()
			cl = strconv.FormatUint(uint64(counter-1), 10)
		})

		It("should return the latest version", func() {
			Expect(response).To(HaveLen(1))
			Expect(response[0].Changelist).To(Equal(cl))
		})
	})

	Context("when version is latest", func() {
		BeforeEach(func() {
			// Counter is latest changelist + 1
			counter := GetChangelistCounter()
			cl = strconv.FormatUint(uint64(counter-1), 10)

			request.Version.Changelist = cl
			request.Source.Filespec.Depot = "..."
			request.Source.Filespec.Stream = "..."
			request.Source.Filespec.Path = "..."
		})

		It("should return an empty list", func() {
			Expect(response).To(HaveLen(0))
		})
	})
})
