// Package english provides utilities for pretty-printing values in English
package english

import "fmt"

// Bool returns either "yes" (true) or "no" (false)
func Bool(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

// OrdinalLiteral returns an ordinal version of the given number.
// E.g.: 1st, 2nd, 123rd, 4916th.
func OrdinalLiteral(n int) string {
	var suffix string

	switch absInt(n) % 10 {
	case 1:
		suffix = "st"
	case 2:
		suffix = "nd"
	case 3:
		suffix = "rd"
	default:
		suffix = "th"
	}

	return fmt.Sprintf("%d%s", n, suffix)
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
