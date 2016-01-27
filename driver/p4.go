package driver

import "os/exec"

func P4(d Driver, args ...string) (string, error) {
	cmd := exec.Command("p4", "-z", "tag", "-p", d.Server.String(), "-u", d.User)
	cmd.Args = append(cmd.Args, args...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
