package upspin

import (
	"bytes"
	"os/exec"
)

type Upspin interface {
	Put(path string, data []byte) error
	Get(path string) ([]byte, error)
	Share(path string) ([]byte, error)
}

func New() Upspin {
	rv := &upspin{}
	return rv
}

type upspin struct{}

func (u *upspin) Put(path string, data []byte) error {
	return nil
}

func (u *upspin) Get(path string) ([]byte, error) {
	cmd := u.createCommand("get", path)
	return u.commandStdOut(cmd)
}

func (u *upspin) Share(path string) ([]byte, error) {
	cmd := u.createCommand("share", "-r", path)
	return u.commandStdOut(cmd)
}

func (u *upspin) commandStdOut(cmd *exec.Cmd) ([]byte, error) {
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func (u *upspin) createCommand(arg ...string) *exec.Cmd {
	return exec.Command("upspin", arg...)
}
