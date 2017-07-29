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
	s := strconv.Itoa(n)
	return fmt.Sprintf("%q", ValueOf(s))
}

// Octal2Binary (n int) int
func Octal2Binary(n int) string {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.OCTAL_BASE, 0)
	utils.CheckError(err, utils.OCTAL, utils.BINARY)
	return decimal.Decimal2Binary(int(i))
}

// Octal2Decimal (n int) string
func Octal2Decimal(n int) string {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.OCTAL_BASE, 0)
	utils.CheckError(err, utils.OCTAL, utils.DECIMAL)
	return strconv.Itoa(int(i))
}

// Octal2Hexadecimal (n int) string
func Octal2Hexadecimal(n int) string {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.OCTAL_BASE, 0)
	utils.CheckError(err, utils.OCTAL, utils.HEXADECIMAL)
	return decimal.Decimal2Hexadecimal(int(i))
}

// String (n int) string
func String(n int) string {
	s := strconv.Itoa(n)
	return fmt.Sprintf("%o", ValueOf(s))
}

// ValueOf (s string) int
func ValueOf(s string) int {
	i, err := strconv.ParseInt(s, utils.OCTAL_BASE, 0)
	utils.CheckError(err, utils.STRING, utils.OCTAL)
	return int(i)
}
