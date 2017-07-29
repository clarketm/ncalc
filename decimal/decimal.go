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

// Decimal2Binary (n int) int
func Decimal2Binary(n int) int {
	s := strconv.FormatInt(int64(n), utils.BINARY_BASE)
	i, err := strconv.Atoi(s)
	utils.CheckError(err, utils.DECIMAL, utils.BINARY)
	return int(i)
}

// Decimal2Octal (n int) int
func Decimal2Octal(n int) int {
	s := strconv.FormatInt(int64(n), utils.OCTAL_BASE)
	i, err := strconv.Atoi(s)
	utils.CheckError(err, utils.DECIMAL, utils.OCTAL)
	return int(i)
}

// Decimal2Hexadecimal (n int) int
func Decimal2Hexadecimal(n int) int {
	s := strconv.FormatInt(int64(n), utils.HEXADECIMAL_BASE)
	i, err := strconv.Atoi(s)
	utils.CheckError(err, utils.DECIMAL, utils.HEXADECIMAL)
	return int(i)
}

// String (n int) string
func String(n int) string {
	return fmt.Sprintf("%d", n)
}
