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
			Host: "localhost",
		}
		user = "Joe_Coder"
		password = ""
		d = new(driver.PerforceDriver)
	})

	Context("when login called", func() {
		Context("with valid username", func() {
			Context("and user does not have password", func() {
				It("should succeed", func() {
					err := d.Login(server, user, password)
					Expect(err).To(Not(HaveOccurred()))
				})
			})
		})
	})
})
