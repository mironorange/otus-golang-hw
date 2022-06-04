package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	var args []string
	name := cmd[0]
	if len(cmd) > 1 {
		args = cmd[1:]
	}

	for envname, envvalue := range env {
		if envvalue.NeedRemove {
			os.Unsetenv(envname)
			continue
		}
		if _, ok := os.LookupEnv(envname); ok {
			os.Unsetenv(envname)
		}
		os.Setenv(envname, envvalue.Value)
	}

	p := exec.Command(name, args...)
	p.Env = os.Environ()
	p.Stdout = os.Stdout
	p.Stdin = os.Stdin
	p.Stderr = os.Stderr

	_ = p.Run()
	return p.ProcessState.ExitCode()
}
