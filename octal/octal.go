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

// Octal2Ascii (n int) string
func Octal2Ascii(n int) string {
	return fmt.Sprintf("%q", Octal2Decimal(n))
}

// Octal2Binary (n int) int
func Octal2Binary(n int) int {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.OCTAL_BASE, 0)
	utils.CheckError(err, utils.OCTAL, utils.BINARY)
	return decimal.Decimal2Binary(int(i))
}

// Octal2Decimal (n int) int
func Octal2Decimal(n int) int {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.OCTAL_BASE, 0)
	utils.CheckError(err, utils.OCTAL, utils.DECIMAL)
	return int(i)
}

// Octal2Hexadecimal (n int) int
func Octal2Hexadecimal(n int) int {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.OCTAL_BASE, 0)
	utils.CheckError(err, utils.OCTAL, utils.HEXADECIMAL)
	return decimal.Decimal2Hexadecimal(int(i))
}

// String (n int) string
func String(n int) string {
	return fmt.Sprintf("%o", Octal2Decimal(n))
}
