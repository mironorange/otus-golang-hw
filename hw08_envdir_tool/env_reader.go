package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	environment := make(Environment)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return environment, err
	}

	for _, info := range files {
		if info.Size() <= 0 {
			environment[info.Name()] = EnvValue{
				NeedRemove: true,
			}
		} else {
			filename := filepath.Join(dir, info.Name())
			if envvalue, err := ReadEnvValueFromFile(filename); err == nil {
				environment[info.Name()] = envvalue
			}
		}
	}

	return environment, nil
}

func ReadEnvValueFromFile(filename string) (EnvValue, error) {
	envvalue := EnvValue{
		NeedRemove: false,
	}

	file, err := os.Open(filename)
	if err != nil {
		return envvalue, err
	}
	defer file.Close()

	br := bufio.NewReader(file)
	line, _, _ := br.ReadLine()
	line = bytes.Replace(line, []byte{0}, []byte{'\n'}, 1)

	linestr := strings.TrimRightFunc(string(line), func(character rune) bool {
		return character == '\n' || character == '\r' || character == ' '
	})
	envvalue.Value = linestr

	return envvalue, nil
}
