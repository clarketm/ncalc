/*

HEXADECIMAL

*/

package hexadecimal

import (
	"fmt"
	"strconv"

	"github.com/clarketm/ncalc/decimal"
	"github.com/clarketm/ncalc/utils"
)

// Hexadecimal2Ascii (n int) string
func Hexadecimal2Ascii(n int) string {
	s := strconv.Itoa(n)
	return fmt.Sprintf("%q", ValueOf(s))
}

// Hexadecimal2Binary (n int) string
func Hexadecimal2Binary(n int) string {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.HEXADECIMAL_BASE, 0)
	utils.CheckError(err, utils.HEXADECIMAL, utils.BINARY)
	return decimal.Decimal2Binary(int(i))
}

// Hexadecimal2Octal (n int) string
func Hexadecimal2Octal(n int) string {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.HEXADECIMAL_BASE, 0)
	utils.CheckError(err, utils.HEXADECIMAL, utils.OCTAL)
	return decimal.Decimal2Octal(int(i))
}

// Hexadecimal2Decimal (n int) string
func Hexadecimal2Decimal(n int) string {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.HEXADECIMAL_BASE, 0)
	utils.CheckError(err, utils.HEXADECIMAL, utils.DECIMAL)
	return strconv.Itoa(int(i))
}

// String (n int) string
func String(n int) string {
	s := strconv.Itoa(n)
	return fmt.Sprintf("%x", ValueOf(s))
}

// ValueOf (s string) int
func ValueOf(s string) int {
	i, err := strconv.ParseInt(s, utils.HEXADECIMAL_BASE, 0)
	utils.CheckError(err, utils.STRING, utils.HEXADECIMAL)
	return int(i)
}
