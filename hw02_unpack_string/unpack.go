package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var (
		out         strings.Builder
		rSymb, rMul rune
		i           int
	)
	rArr := []rune(s)
	for {
		switch {
		case (len(rArr) - i) <= 0:
			return out.String(), nil
		case (len(rArr) - i) == 1:
			rSymb = rArr[i]
			rMul = 0
		default:
			rSymb = rArr[i]
			rMul = rArr[i+1]
		}
		if unicode.IsDigit(rSymb) {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(rMul) {
			mul, err := strconv.Atoi(string(rMul))
			if err != nil {
				return "", err
			}
			if mul != 0 {
				out.WriteString(strings.Repeat(string(rSymb), mul))
			}
			i += 2
		} else {
			out.WriteRune(rSymb)
			i++
		}
	}
}
