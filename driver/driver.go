package driver

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/olhtbr/p4-resource/models"
)

type Driver interface {
	Login(models.Server, string, string) error
	GetLatestChangelist(models.Filespec) (string, error)
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

func (d PerforceDriver) GetLatestChangelist(f models.Filespec) (string, error) {
	cmd := exec.Command("p4", "-z", "tag", "-p", fmt.Sprintf("%s", d.server), "-u", d.user, "changes", "-m", "1", fmt.Sprintf("%s", f))
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Empty output can be returned by p4, eg. on non-existent filespec
	if len(out) == 0 {
		return "", nil
	}

	re := regexp.MustCompile(`\.\.\.\s+change\s+(\d+)`)
	if match := re.FindStringSubmatch(string(out)); len(match) == 2 {
		return match[1], nil
	}

	return "", nil
}

func (d PerforceDriver) GetTicket() string {
	return d.ticket
}
