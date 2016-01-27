package driver_test

import (
	"github.com/olhtbr/p4-resource/driver"
	"github.com/olhtbr/p4-resource/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Driver", func() {
	var d *driver.Driver
	var server models.Server
	var user string
	var password string
	var f models.Filespec

	BeforeEach(func() {
		server = models.Server{
			Host: "localhost",
		}
		user = "Joe_Coder"
		password = ""
		f = models.Filespec{
			Depot:  "...",
			Stream: "...",
			Path:   "...",
		}
		d = new(driver.Driver)
	})

	Context("when login called", func() {
		Context("with valid username", func() {
			Context("and user does not have password", func() {
				It("should succeed", func() {
					err := d.Login(server, user, password)
					Expect(err).To(Not(HaveOccurred()))
					Expect(d.Server).To(Equal(server))
					Expect(d.User).To(Equal(user))
				})
			})
		})

		Context("with invalid username", func() {
			BeforeEach(func() {
				user = "non-existent-user"
			})

			It("should fail", func() {
				err := d.Login(server, user, password)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Context("when latest changelist is requested", func() {
		It("should return it", func() {
			cl, err := d.GetLatestChangelist(f)
			Expect(err).To(Not(HaveOccurred()))
			Expect(cl).To(Equal("12104"))
		})
	})

	Context("when changelists newer than a specific changelist are requested", func() {
		Context("and the specified changelist is submitted", func() {
			It("should not return the specfied changelist", func() {
				cls, err := d.GetChangelistsNewerThan("12099", f)
				Expect(err).To(Not(HaveOccurred()))
				Expect(cls).To(Not(ContainElement("12099")))
			})

			It("should return the correct list", func() {
				cls, err := d.GetChangelistsNewerThan("12099", f)
				Expect(err).To(Not(HaveOccurred()))
				Expect(cls).To(HaveLen(2))
				Expect(cls[0]).To(Equal("12103"))
				Expect(cls[1]).To(Equal("12104"))
			})
		})

		Context("and the specified changelist does not exist yet", func() {
			It("should return and empty list", func() {
				cls, err := d.GetChangelistsNewerThan("15000", f)
				Expect(err).To(Not(HaveOccurred()))
				Expect(cls).To(BeEmpty())
			})
		})

		Context("and the specified changelist is pending", func() {
			It("should not return the specfied changelist", func() {
				cls, err := d.GetChangelistsNewerThan("12100", f)
				Expect(err).To(Not(HaveOccurred()))
				Expect(cls).To(Not(ContainElement("12100")))
			})

			It("should return the correct list", func() {
				cls, err := d.GetChangelistsNewerThan("12100", f)
				Expect(err).To(Not(HaveOccurred()))
				Expect(cls).To(HaveLen(2))
				Expect(cls[0]).To(Equal("12103"))
				Expect(cls[1]).To(Equal("12104"))
			})
		})
	})

	Context("when a changelist does not exist", func() {
		It("should find it out", func() {
			exists, err := d.ChangelistExists("15000")
			Expect(err).To(Not(HaveOccurred()))
			Expect(exists).To(BeFalse())
		})
	})

	Context("when a changelist exists", func() {
		Context("and is submitted", func() {
			It("should find it out", func() {
				exists, err := d.ChangelistExists("12099")
				Expect(err).To(Not(HaveOccurred()))
				Expect(exists).To(BeTrue())
			})
		})

		Context("and is pending", func() {
			It("should find it out", func() {
				exists, err := d.ChangelistExists("12100")
				Expect(err).To(Not(HaveOccurred()))
				Expect(exists).To(BeTrue())
			})
		})
	})
})
