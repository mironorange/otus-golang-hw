package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmoji(t *testing.T) {
	input := "Девушка🔥3"
	result, _ := Unpack(input)
	require.Equal(t, result, "Девушка🔥🔥🔥")
}

func TestCyrillic(t *testing.T) {
	input := "П1р3вет4"
	result, _ := Unpack(input)
	require.Equal(t, result, "Прррветттт")
}

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: `abcd`, expected: `abcd`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestEmptyStringLength(t *testing.T) {
	input := ""
	result, _ := Unpack(input)
	require.Equal(t, 0, len(result))
}

func TestUnpackWithSpecialSymbols(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: `d\n5abc`, expected: `d\n\n\n\n\nabc`},
		{input: `dn5abc`, expected: `dnnnnnabc`},
		{input: `dn5\abc`, expected: `dnnnnnabc`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestErrorReturning(t *testing.T) {
	input := "111"
	_, err := Unpack(input)
	require.Error(t, err)
}
