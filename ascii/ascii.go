/*

ASCII

*/

package ascii

import (
	"fmt"

	"github.com/clarketm/ncalc/decimal"
)

// Ascii2Binary (s string) int
func Ascii2Binary(n rune) int {
	return decimal.Decimal2Binary(int(n))
}

// Ascii2Octal (s string) int
func Ascii2Octal(n rune) int {
	return decimal.Decimal2Octal(int(n))
}

// Ascii2Decimal (c string) int
func Ascii2Decimal(n rune) int {
	return int(n)
}

// Ascii2Hexadecimal (c string) int
func Ascii2Hexadecimal(n rune) int {
	return decimal.Decimal2Hexadecimal(int(n))
}

// String (n string) string
func String(n rune) string {
	return fmt.Sprintf("%c", n)
}
