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
	shielding := false // Shielding mode
	sequences := make([]byte, 0)
	for _, char := range src {
		if unicode.IsDigit(char) && !shielding {
			if len(sc) <= 0 {
				return "", ErrInvalidString
			}
			ml, err := strconv.Atoi(string(char))
			if err != nil {
				return "", ErrInvalidString
			}

			sequences = append(sequences, bytes.Repeat(sc, ml)...)
			ml = 1
			sc = sc[:0]
		} else {
			if char == '\\' && !shielding {
				if len(sc) > 0 {
					sequences = append(sequences, sc...)
					ml = 1
					sc = sc[:0]
				}

				shielding = true
				continue
			}
			if len(sc) > 0 {
				sequences = append(sequences, sc...)
				ml = 1
				sc = sc[:0]
			}

			sc = append(sc, byte(char))
			shielding = false
		}
	}

	if len(sc) > 0 {
		sequences = append(sequences, bytes.Repeat(sc, ml)...)
	}

	return string(sequences), nil
}
