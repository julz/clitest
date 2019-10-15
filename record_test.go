package clitest_test

import (
	"os/exec"
	"testing"

	"github.com/julz/clitest"
	"gotest.tools/assert"
)

func TestRun(t *testing.T) {
	examples := []struct {
		Title  string
		Cmd    *exec.Cmd
		Expect clitest.Result
	}{
		{
			Title: "Stdout",
			Cmd:   exec.Command("echo", "hello"),
			Expect: clitest.Result{
				Stdout: "hello\n",
			},
		},
		{
			Title: "Stderr",
			Cmd:   exec.Command("bash", "-c", "echo hi >&2"),
			Expect: clitest.Result{
				Stderr: "hi\n",
			},
		},
		{
			Title: "Sh helper",
			Cmd:   clitest.Sh("echo 'stderr via sh' >&2"),
			Expect: clitest.Result{
				Stderr: "stderr via sh\n",
			},
		},
		{
			Title: "ExitCode",
			Cmd:   clitest.Sh("echo 'stdout'; echo 'stderr via sh' >&2; exit 12"),
			Expect: clitest.Result{
				Stdout:   "stdout\n",
				Stderr:   "stderr via sh\n",
				ExitCode: 12,
			},
		},
	}

	for _, eg := range examples {
		t.Run(eg.Title, func(t *testing.T) {
			result := clitest.Run(eg.Cmd)
			assert.DeepEqual(t, eg.Expect, result)
		})
	}
}
