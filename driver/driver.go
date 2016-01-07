package driver

import "github.com/olhtbr/p4-resource/models"

type Driver interface {
	GetLatestChangelist(models.Filespec) string
}

type PerforceDriver struct {
	Server models.Server
	User   string
	Ticket string
}

func (d PerforceDriver) GetLatestChangelist(f models.Filespec) string {
	return "123456"
}
