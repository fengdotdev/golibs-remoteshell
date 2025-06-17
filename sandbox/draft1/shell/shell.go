package shell

import (
	"bytes"
	"os/exec"
)

type Shell struct{}

func New() *Shell {
	return &Shell{}
}

func (s *Shell) RunCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out // Capture stderr in case of errors
	err := cmd.Run()
	if err != nil {
		return out.String(), err
	}
	return out.String(), nil
}
