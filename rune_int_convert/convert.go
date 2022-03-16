package rune_int_convert

import (
	"regexp"
)

// ConvertRune converts rune A=0, B=1, works for [A...Z] if not in interval return -1
func ConvertRune(r rune) int {
	i := int(r)

	if i < 65 || i > 90 {
		return -1
	}
	return i - 64
}

var reg = regexp.MustCompile("^[A-Z]+$")

// ConvertString convert string to int AA = 27, AB = 28 etc. , if one of runes not in [A..Z] returns -1
func ConvertString(s string) int {
	if !reg.Match([]byte(s)) {
		return -1
	}

	l := len(s)

	runes := []rune(s)

	var sum int
	var pow = 1
	for i := l - 1; i >= 0; i-- {
		sum += (int(runes[i]) - 64) * pow
		pow *= 26
	}

	return sum
}
