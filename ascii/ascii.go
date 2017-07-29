/*

ASCII

*/

package ascii

import (
	"fmt"

	"github.com/clarketm/ncalc/decimal"
)

// Ascii2Binary (s rune) string
func Ascii2Binary(n rune) string {
	return decimal.Decimal2Binary(int(n))
}

// Ascii2Octal (s rune) string
func Ascii2Octal(n rune) string {
	return decimal.Decimal2Octal(int(n))
}

// Ascii2Decimal (c rune) string
func Ascii2Decimal(n rune) string {
	return decimal.String(int(n))
}

// Ascii2Hexadecimal (c rune) string
func Ascii2Hexadecimal(n rune) string {
	return decimal.Decimal2Hexadecimal(int(n))
}

// String (n rune) string
func String(n rune) string {
	return fmt.Sprintf("%q", n)
}

// ValueOf (s string) rune
// TODO
