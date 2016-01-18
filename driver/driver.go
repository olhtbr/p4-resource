package driver

import (
	"os/exec"
	"regexp"
	"sort"
	"strconv"

	"github.com/olhtbr/p4-resource/models"
)

type Driver interface {
	Login(models.Server, string, string) error
	GetLatestChangelist(models.Filespec) (string, error)
	GetChangelistsNewerThan(string, models.Filespec) ([]string, error)
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

	cmd := exec.Command("p4", "-p", server.String(), "-u", user, "-P", password, "login", "-p")
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

func (d PerforceDriver) GetLatestChangelist(f models.Filespec) (string, error) {
	cmd := exec.Command("p4", "-z", "tag", "-p", d.server.String(), "-u", d.user, "changes", "-m", "1", f.String())
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

func (d PerforceDriver) GetChangelistsNewerThan(cl string, f models.Filespec) ([]string, error) {
	clNum, err := strconv.ParseUint(cl, 10, 64)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("p4", "-z", "tag", "-p", d.server.String(), "-u", d.user, "changes", f.String()+"@"+strconv.FormatUint(clNum+1, 10)+",#head")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	cls := []string{}
	if len(out) == 0 {
		return cls, nil
	}

	re := regexp.MustCompile(`\.\.\.\s+change\s+(\d+)`)
	matches := re.FindAllStringSubmatch(string(out), -1)
	for _, match := range matches {
		cls = append(cls, match[1])
	}

	sort.Strings(cls)

	return cls, nil
}

func (d PerforceDriver) GetTicket() string {
	return d.ticket
}
