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

// Binary2Ascii (n int) string
func Binary2Ascii(n int) string {
	s := strconv.Itoa(n)
	return fmt.Sprintf("%q", ValueOf(s))
}

// Binary2Octal (n int) string
func Binary2Octal(n int) string {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.BINARY_BASE, 0)
	utils.CheckError(err, utils.BINARY, utils.OCTAL)
	return decimal.Decimal2Octal(int(i))
}

// Binary2Decimal (n int) string
func Binary2Decimal(n int) string {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.BINARY_BASE, 0)
	utils.CheckError(err, utils.BINARY, utils.DECIMAL)
	return strconv.Itoa(int(i))
}

// Binary2Hexadecimal (n int) string
func Binary2Hexadecimal(n int) string {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.BINARY_BASE, 0)
	utils.CheckError(err, utils.BINARY, utils.HEXADECIMAL)
	return decimal.Decimal2Hexadecimal(int(i))
}

// String (n int) string
func String(n int) string {
	s := strconv.Itoa(n)
	return fmt.Sprintf("%b", ValueOf(s))
}

// ValueOf (s string) int
func ValueOf(s string) int {
	i, err := strconv.ParseInt(s, utils.BINARY_BASE, 0)
	utils.CheckError(err, utils.STRING, utils.BINARY)
	return int(i)
}
