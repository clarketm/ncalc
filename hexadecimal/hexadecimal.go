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
	return fmt.Sprintf("%q", Hexadecimal2Decimal(n))
}

// Hexadecimal2Binary (n int) int
func Hexadecimal2Binary(n int) int {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.HEXADECIMAL_BASE, 0)
	utils.CheckError(err, utils.HEXADECIMAL, utils.BINARY)
	return decimal.Decimal2Binary(int(i))
}

// Hexadecimal2Octal (n int) int
func Hexadecimal2Octal(n int) int {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.HEXADECIMAL_BASE, 0)
	utils.CheckError(err, utils.HEXADECIMAL, utils.OCTAL)
	return decimal.Decimal2Octal(int(i))
}

// Hexadecimal2Decimal (n int) int
func Hexadecimal2Decimal(n int) int {
	s := strconv.Itoa(n)
	i, err := strconv.ParseInt(s, utils.HEXADECIMAL_BASE, 0)
	utils.CheckError(err, utils.HEXADECIMAL, utils.DECIMAL)
	return int(i)
}

// String (n int) string
func String(n int) string {
	return fmt.Sprintf("%x", Hexadecimal2Decimal(n))
}
