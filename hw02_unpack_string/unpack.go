package hw02unpackstring

import (
	"bytes"
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(src string) (string, error) {
	if len(src) <= 0 {
		return "", nil
	}

	sc := make([]byte, 0)
	ml := 1
	sequences := make([]byte, 0)
	for _, char := range src {
		if unicode.IsDigit(char) {
			ml, err := strconv.Atoi(string(char))
			if err != nil {
				return "", ErrInvalidString
			}
			sequences = append(sequences, bytes.Repeat(sc, ml)...)
			ml = 1
			sc = sc[:0]
		} else {
			if len(sc) > 0 {
				sequences = append(sequences, sc...)
				ml = 1
				sc = sc[:0]
			}
			sc = append(sc, byte(char))
		}
	}

	if len(sc) > 0 {
		sequences = append(sequences, bytes.Repeat(sc, ml)...)
	} else {
		return "", ErrInvalidString
	}

	return string(sequences), nil
}
