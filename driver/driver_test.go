package driver_test

import (
	"github.com/olhtbr/p4-resource/driver"
	"github.com/olhtbr/p4-resource/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Driver", func() {
	var d driver.Driver
	var server models.Server
	var user string
	var password string

	BeforeEach(func() {
		server = models.Server{
			Protocol: "",
			Host:     "localhost",
			Port:     1666,
		}
		user = "Joe_Coder"
		password = ""
		d = driver.PerforceDriver{}
	})

	Context("when login called", func() {
		Context("with valid username and password", func() {
			It("should store a valid ticket", func() {
				err := d.Login(server, user, password)
				Expect(err).To(Not(HaveOccurred()))
				Expect(d.GetTicket()).To(Not(BeNil()))
				Expect(d.GetTicket()).To(Not(BeEmpty()))
			})
		})
	})
})
