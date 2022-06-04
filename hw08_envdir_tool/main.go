package main

import (
	"fmt"
	"os"
)

const (
	ErrInvalid  = "invalid argument"
	ErrNotExist = "directory does not exist"
)

func main() {
	dir, cmd, err := prepareArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	environment, err := ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(RunCmd(cmd, environment))
}

func prepareArgs(args []string) (string, []string, error) {
	if len(args) < 3 {
		return "", nil, fmt.Errorf(ErrInvalid)
	}

	dir := args[1]
	if stat, err := os.Stat(dir); err != nil || !stat.IsDir() {
		return "", nil, fmt.Errorf(ErrNotExist)
	}

	return dir, args[2:], nil
}
