package driver

import (
	"fmt"
	"os/exec"

	"github.com/olhtbr/p4-resource/models"
)

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

func (d *PerforceDriver) Login(server models.Server, user string, password string) error {
	d.server = server
	d.user = user

	cmd := exec.Command("p4", "-p", fmt.Sprintf("%s", server), "-u", user, "-P", password, "login", "-p")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func (d PerforceDriver) GetLatestChangelist(f models.Filespec) string {
	return "123456"
}

func (d PerforceDriver) GetTicket() string {
	return d.ticket
}
