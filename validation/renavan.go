package validation

import (
	"strconv"
)

var (
	renavamAcc = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

// IsRenavanValid verifies if the given string is a valid RENAVAM document.
func IsRenavanValid(renavan string) bool {
	if len(renavan) != 11 {
		return false
	}
	if !allDigit(renavan) {
		return false
	}

	var sum int
	for i, r := range renavan[:len(renavan)-1] {
		sum += toInt(r) * renavamAcc[i]
	}

	digit := (sum * 10) % 11
	if digit == 10 {
		digit = 0
	}

	return string(renavan[len(renavan)-1]) == strconv.Itoa(digit)
}