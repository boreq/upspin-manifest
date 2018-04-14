// Package upspin handles upspin IO. Currently this is done by directly calling
// the upspin executable.
package upspin

import (
	"bytes"
	"fmt"
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
	var stdErr bytes.Buffer
	in := bytes.NewBuffer(data)

	cmd := u.createCommand("put", path)
	cmd.Stdin = in
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%s: %s", err, stdErr.Bytes())
	}
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
	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", err, stdErr.Bytes())
	}
	return stdOut.Bytes(), nil
}

func (u *upspin) createCommand(arg ...string) *exec.Cmd {
	return exec.Command("upspin", arg...)
}
