package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDir(t *testing.T) {
	expected := Environment{
		"BAR": EnvValue{
			Value:      "bar",
			NeedRemove: false,
		},
		"EMPTY": EnvValue{
			Value:      "",
			NeedRemove: false,
		},
		"FOO": EnvValue{
			Value:      "   foo\nwith new line",
			NeedRemove: false,
		},
		"HELLO": EnvValue{
			Value:      "\"hello\"",
			NeedRemove: false,
		},
		"UNSET": EnvValue{
			NeedRemove: true,
		},
	}
	result, err := ReadDir("testdata/env")
	assert.Equal(t, expected, result)
	assert.NoError(t, err)

	nonexistentdir := "testdata/nonexistentdirectory"
	assert.NoDirExists(t, nonexistentdir)
	_, err = ReadDir(nonexistentdir)
	assert.ErrorContains(t, err, "no such file or directory")
}
