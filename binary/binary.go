/*

BINARY

*/

package binary

import (
	"fmt"
	"strconv"

	"github.com/clarketm/ncalc/decimal"
	"github.com/clarketm/ncalc/utils"
)

// Binary2Ascii (s string) string
func Binary2Ascii(s string) string {
	return fmt.Sprintf("%q", ValueOf(s))
}

// Binary2Octal (s string) string
func Binary2Octal(s string) string {
	i, err := utils.Parse(s, utils.BINARY_BASE)
	utils.CheckError(err, utils.BINARY, utils.OCTAL)
	s = strconv.Itoa(int(i))
	return decimal.Decimal2Octal(s)
}

// Binary2Decimal (s string) string
func Binary2Decimal(s string) string {
	i, err := utils.Parse(s, utils.BINARY_BASE)
	utils.CheckError(err, utils.BINARY, utils.DECIMAL)
	return strconv.Itoa(int(i))
}

// Binary2Hexadecimal (s string) string
func Binary2Hexadecimal(s string) string {
	i, err := utils.Parse(s, utils.BINARY_BASE)
	utils.CheckError(err, utils.BINARY, utils.HEXADECIMAL)
	s = strconv.Itoa(int(i))
	return decimal.Decimal2Hexadecimal(s)
}

// String (s string) string
func String(s string) string {
	return fmt.Sprintf("%b", ValueOf(s))
}

// ValueOf (s string) int
func ValueOf(s string) int {
	i, err := utils.Parse(s, utils.BINARY_BASE)
	utils.CheckError(err, utils.STRING, utils.BINARY)
	return int(i)
}
