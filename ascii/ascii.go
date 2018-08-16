/*

ASCII

*/

package ascii

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/clarketm/ncalc/decimal"
)

// Ascii2Binary (s string) string
func Ascii2Binary(s string) string {
	s = strconv.Itoa(ValueOf(s))
	return decimal.Decimal2Binary(s)
}

// Ascii2Octal (s string) string
func Ascii2Octal(s string) string {
	s = strconv.Itoa(ValueOf(s))
	return decimal.Decimal2Octal(s)
}

// Ascii2Decimal (s string) string
func Ascii2Decimal(s string) string {
	s = strconv.Itoa(ValueOf(s))
	return decimal.String(s)
}

// Ascii2Hexadecimal (s string) string
func Ascii2Hexadecimal(s string) string {
	s = strconv.Itoa(ValueOf(s))
	return decimal.Decimal2Hexadecimal(s)
}

// String (s string) string
func String(s string) string {
	return fmt.Sprintf("%q", ValueOf(s))
}

// ValueOf (s string) int
func ValueOf(s string) int {
	// if utils.IsLiteral(s) {
	// 	ss, _ := strconv.ParseInt(s, 0, 0)
	// }
	s, _ = strconv.Unquote(`"` + s + `"`)
	r, _ := utf8.DecodeRuneInString(s)
	return int(r)
}

// ValueOf (s string) rune
// TODO
