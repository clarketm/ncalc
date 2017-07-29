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
	return fmt.Sprintf("%q", Binary2Decimal(n))
}

// Binary2Octal (n int) int
func Binary2Octal(n int) int {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.BINARY_BASE, 0)
	utils.CheckError(err, utils.BINARY, utils.OCTAL)
	return decimal.Decimal2Octal(int(i))
}

// Binary2Decimal (n int) int
func Binary2Decimal(n int) int {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.BINARY_BASE, 0)
	utils.CheckError(err, utils.BINARY, utils.DECIMAL)
	return int(i)
}

// Binary2Hexadecimal (n int) int
func Binary2Hexadecimal(n int) int {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.BINARY_BASE, 0)
	utils.CheckError(err, utils.BINARY, utils.HEXADECIMAL)
	return decimal.Decimal2Hexadecimal(int(i))
}

// String (n int) string
func String(n int) string {
	return fmt.Sprintf("%b", Binary2Decimal(n))
}
