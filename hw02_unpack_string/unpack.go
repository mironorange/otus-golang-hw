package hw02unpackstring

import (
	"bytes"
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(src string) (string, error) {
	sequences := make([]byte, 0)
	sc := make([]byte, 0) // sequence
	sequence := make([]byte, 0)
	ml := make([]byte, 0) // multiplier
	for _, char := range src {
		if unicode.IsDigit(char) {
			if len(sc) > 0 {
				sequence = sc
				sc = sc[:0]
			}

			ml = append(ml, byte(char))
		} else {
			if len(sequence) > 0 {
				if len(ml) > 0 {
					multiplier, err := strconv.Atoi(string(ml))
					if err != nil {
						return "", ErrInvalidString
					}
					ml = ml[:0]
					sequences = append(sequences, bytes.Repeat(sequence, multiplier)...)
				} else {
					sequences = append(sequences, sequence...)
				}
			}

			sc = append(sc, byte(char))
		}
	}

	if len(sequence) > 0 {
		if len(ml) > 0 {
			multiplier, err := strconv.Atoi(string(ml))
			if err != nil {
				return "", ErrInvalidString
			}
			sequences = append(sequences, bytes.Repeat(sequence, multiplier)...)
		} else {
			sequences = append(sequences, sequence...)
		}
	} else {
		return "", ErrInvalidString
	}

	return string(sequences), nil
}
