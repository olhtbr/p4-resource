package driver

import "github.com/olhtbr/p4-resource/models"

type Driver interface {
	Login(models.Server, string, string) error
	GetLatestChangelist(models.Filespec) string
	GetTicket() string
}

type PerforceDriver struct {
	server models.Server
	user   string
	ticket string
}

func (d *PerforceDriver) Login(s models.Server, u string, p string) error {
	d.server = s
	d.user = u

	return nil
}

func (d PerforceDriver) GetLatestChangelist(f models.Filespec) string {
	return "123456"
}

func (d PerforceDriver) GetTicket() string {
	return d.ticket
}
