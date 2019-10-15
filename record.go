package clitest

import (
	"bytes"
	"os/exec"
)

type Result struct {
	ExitCode int
	Stdout   string
	Stderr   string
}

// Run runs the cmd and returns a struct with the result for easy assertion.
//
// The passed command must be valid (e.g. exec.LookPath must succeed) or this
// method will panic.
func Run(cmd *exec.Cmd) Result {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	status := 0
	if err != nil {
		eerr, ok := err.(*exec.ExitError)

		if !ok {
			// we were given a cmd to run that doesn't exist
			panic(err)
		}

		status = eerr.ProcessState.ExitCode()
	}

	return Result{
		Stdout:   string(stdout.Bytes()),
		Stderr:   string(stderr.Bytes()),
		ExitCode: status,
	}
}

// Sh returns an exec.Cmd that runs the given string in a shell.
// (i.e. it is a tiny bit of sugar around just doing `exec.Command("sh", "-c", cmd)`)
func Sh(cmd string) *exec.Cmd {
	return exec.Command("sh", "-c", cmd)
}
