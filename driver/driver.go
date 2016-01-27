package driver

import (
	"os/exec"
	"regexp"
	"sort"
	"strconv"

	"github.com/olhtbr/p4-resource/models"
)

type Driver struct {
	Server models.Server
	User   string
	Ticket string
}

func (d *Driver) Login(server models.Server, user string, password string) error {
	d.Server = server
	d.User = user

	cmd := exec.Command("p4", "-p", server.String(), "-u", user, "-P", password, "login", "-p")
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

func (d Driver) GetLatestChangelist(f models.Filespec) (string, error) {
	cmd := exec.Command("p4", "-z", "tag", "-p", d.Server.String(), "-u", d.User, "changes", "-m", "1", f.String())
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

func (d Driver) GetChangelistsNewerThan(cl string, f models.Filespec) ([]string, error) {
	clNum, err := strconv.ParseUint(cl, 10, 64)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("p4", "-z", "tag", "-p", d.Server.String(), "-u", d.User, "changes", f.String()+"@"+strconv.FormatUint(clNum+1, 10)+",#head")
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

func (d Driver) ChangelistExists(cl string) (bool, error) {
	cmd := exec.Command("p4", "-z", "tag", "-p", d.Server.String(), "-u", d.User, "describe", cl)
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}

	matched, err := regexp.Match("no such changelist", out)
	if err != nil {
		return false, err
	}

	if len(out) == 0 || matched {
		return false, nil
	}

	return true, nil
}
