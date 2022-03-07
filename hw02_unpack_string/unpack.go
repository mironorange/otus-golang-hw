package hw02unpackstring

import (
	"bytes"
	"errors"
	"strconv"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
	SpecialSymbols   = map[rune]int{
		'n': 1,
		'r': 1,
		't': 1,
	}
)

func Unpack(src string) (string, error) {
	if len(src) == 0 {
		return "", nil
	}

	sc := make([]byte, 0)
	ml := 1
	shielding := false // Shielding mode
	sequences := make([]byte, 0)
	for _, char := range src {
		// Если повстречалась цифра или включен режим экранирования
		if unicode.IsDigit(char) && !shielding {
			if len(sc) == 0 {
				return "", ErrInvalidString
			}
			ml, err := strconv.Atoi(string(char))
			if err != nil {
				return "", ErrInvalidString
			}
			sequences = append(sequences, bytes.Repeat(sc, ml)...)
			sc = sc[:0]
			continue
		}
		// Если встретился обратный слеш, то проигнорировать символ и включить режим экранирования
		if char == '\\' && !shielding {
			if len(sc) > 0 {
				sequences = append(sequences, sc...)
				ml = 1
				sc = sc[:0]
			}
			shielding = true
			continue
		}
		// В любом ином случае рассматриваем как символ, который нужно распаковать
		if len(sc) > 0 {
			sequences = append(sequences, sc...)
			ml = 1
			sc = sc[:0]
		}
		if _, ok := SpecialSymbols[char]; ok && shielding {
			sc = append(sc, '\\')
		}
		sc = append(sc, byte(char))
		shielding = false
	}

	if len(sc) > 0 {
		sequences = append(sequences, bytes.Repeat(sc, ml)...)
	}

	return string(sequences), nil
}
