package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
	SpecialSymbols   = map[rune]struct{}{
		'n': {},
		'r': {},
		't': {},
	}
)

func Unpack(src string) (string, error) {
	if len(src) == 0 {
		return "", nil
	}

	var s strings.Builder

	m := false // Shielding mode
	sc := make([]rune, 0)
	ml := 1

	for _, char := range src {
		// Если повстречалась цифра или включен режим экранирования
		if unicode.IsDigit(char) && !m {
			if len(sc) == 0 {
				return "", ErrInvalidString
			}
			ml, err := strconv.Atoi(string(char))
			if err != nil {
				return "", ErrInvalidString
			}
			s.WriteString(strings.Repeat(string(sc), ml))
			sc = sc[:0]
			continue
		}
		// Если встретился обратный слеш, то проигнорировать символ и включить режим экранирования
		if char == '\\' && !m {
			if len(sc) > 0 {
				s.WriteString(string(sc))
				ml = 1
				sc = sc[:0]
			}
			m = true
			continue
		}
		// В любом ином случае рассматриваем как символ, который нужно распаковать
		if len(sc) > 0 {
			s.WriteString(string(sc))
			ml = 1
			sc = sc[:0]
		}
		if _, ok := SpecialSymbols[char]; ok && m {
			sc = append(sc, '\\')
		}
		sc = append(sc, char)
		m = false
	}

	if len(sc) > 0 {
		s.WriteString(strings.Repeat(string(sc), ml))
	}

	return s.String(), nil
}
