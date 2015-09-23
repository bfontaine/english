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

	sign, n := absInt(n)

	if d := n % 100; d > 3 && d <= 20 {
		suffix = "th"
	} else {
		switch n % 10 {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		case 3:
			suffix = "rd"
		default:
			suffix = "th"
		}
	}

	return fmt.Sprintf("%s%d%s", sign, n, suffix)
}

// Plural is a simple utility which takes a word and a quantity, and returns
// the pluralized version of the word if necessary. It doesn’t support special
// cases (e.g. Plural("woman", 2) will return "womans") but makes code more
// readable in the general case.
func Plural(s string, n int) string {
	if _, n = absInt(n); n != 1 {
		return s + "s"
	}

	return s
}

func absInt(n int) (string, int) {
	if n < 0 {
		return "-", -n
	}
	return "", n
}
