/*

DECIMAL

*/

package decimal

import (
	"fmt"
	"strconv"

	"github.com/clarketm/ncalc/utils"
)

// Decimal2Ascii (n int) string
func Decimal2Ascii(n int) string {
	return fmt.Sprintf("%q", n)
}

// Decimal2Binary (n int) string
func Decimal2Binary(n int) string {
	s := strconv.FormatInt(int64(n), utils.BINARY_BASE)
	return s
}

// Decimal2Octal (n int) string
func Decimal2Octal(n int) string {
	s := strconv.FormatInt(int64(n), utils.OCTAL_BASE)
	return s
}

// Decimal2Hexadecimal (n int) string
func Decimal2Hexadecimal(n int) string {
	s := strconv.FormatInt(int64(n), utils.HEXADECIMAL_BASE)
	return s
}

// String (n int) string
func String(n int) string {
	return fmt.Sprintf("%d", n)
}

// ValueOf (s string) int
func ValueOf(s string) int {
	i, err := strconv.Atoi(s)
	utils.CheckError(err, utils.STRING, utils.DECIMAL)
	return int(i)
}
