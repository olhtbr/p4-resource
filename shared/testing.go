package shared

import (
	"encoding/json"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func Run(cmd *exec.Cmd, request interface{}, response interface{}, code *int) {
	JustBeforeEach(func() {
		stdin, err := cmd.StdinPipe()
		Expect(err).To(Not(HaveOccurred()))

		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).To(Not(HaveOccurred()))

		err = json.NewEncoder(stdin).Encode(request)
		Expect(err).To(Not(HaveOccurred()))

		Eventually(session).Should(gexec.Exit(*code))

		if *code == 0 {
			err = json.Unmarshal(session.Out.Contents(), response)
			Expect(err).To(Not(HaveOccurred()))
		}
	})
}
