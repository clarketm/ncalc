/*

OCTAL

*/

package octal

import (
	"fmt"
	"strconv"

	"github.com/clarketm/ncalc/decimal"
	"github.com/clarketm/ncalc/utils"
)

// Octal2Ascii (s string) string
func Octal2Ascii(s string) string {
	return fmt.Sprintf("%q", ValueOf(s))
}

// Octal2Binary (s string) string
func Octal2Binary(s string) string {
	i, err := utils.Parse(s, utils.OCTAL_BASE)
	utils.CheckError(err, utils.OCTAL, utils.BINARY)
	s = strconv.Itoa(int(i))
	return decimal.Decimal2Binary(s)
}

// Octal2Decimal (s string) string
func Octal2Decimal(s string) string {
	i, err := utils.Parse(s, utils.OCTAL_BASE)
	utils.CheckError(err, utils.OCTAL, utils.DECIMAL)
	return strconv.Itoa(int(i))
}

// Octal2Hexadecimal (s string) string
func Octal2Hexadecimal(s string) string {
	i, err := utils.Parse(s, utils.OCTAL_BASE)
	utils.CheckError(err, utils.OCTAL, utils.HEXADECIMAL)
	s = strconv.Itoa(int(i))
	return decimal.Decimal2Hexadecimal(s)
}

// String (s string) string
func String(s string) string {
	return fmt.Sprintf("%o", ValueOf(s))
}

// ValueOf (s string) int
func ValueOf(s string) int {
	i, err := utils.Parse(s, utils.OCTAL_BASE)
	utils.CheckError(err, utils.STRING, utils.OCTAL)
	return int(i)
}
