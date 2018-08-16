/*

DECIMAL

*/

package decimal

import (
	"fmt"
	"strconv"

	"github.com/clarketm/ncalc/utils"
)

// Decimal2Ascii (s string) string
func Decimal2Ascii(s string) string {
	return fmt.Sprintf("%q", ValueOf(s))
}

// Decimal2Binary (s string) string
func Decimal2Binary(s string) string {
	i, err := utils.Parse(s, utils.DECIMAL_BASE)
	utils.CheckError(err, utils.STRING, utils.DECIMAL)
	s = strconv.FormatInt(int64(i), utils.BINARY_BASE)
	return s
}

// Decimal2Octal (s string) string
func Decimal2Octal(s string) string {
	i, err := utils.Parse(s, utils.DECIMAL_BASE)
	utils.CheckError(err, utils.STRING, utils.DECIMAL)
	s = strconv.FormatInt(int64(i), utils.OCTAL_BASE)
	return s
}

// Decimal2Hexadecimal (s string) string
func Decimal2Hexadecimal(s string) string {
	i, err := utils.Parse(s, utils.DECIMAL_BASE)
	utils.CheckError(err, utils.STRING, utils.DECIMAL)
	s = strconv.FormatInt(int64(i), utils.HEXADECIMAL_BASE)
	return s
}

// String (s string) string
func String(s string) string {
	return fmt.Sprintf("%d", ValueOf(s))
}

// ValueOf (s string) int
func ValueOf(s string) int {
	i, err := utils.Parse(s, utils.DECIMAL_BASE)
	utils.CheckError(err, utils.STRING, utils.DECIMAL)
	return int(i)
}
