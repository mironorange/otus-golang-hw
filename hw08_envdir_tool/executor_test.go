package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCmd(t *testing.T) {
	var exitcode int

	exitcode = RunCmd([]string{"echo", "\"Hello, World\""}, Environment{})
	assert.Equal(t, 0, exitcode)

	exitcode = RunCmd([]string{"testdata/exitcode.sh", "5"}, Environment{})
	assert.Equal(t, 5, exitcode)
}
