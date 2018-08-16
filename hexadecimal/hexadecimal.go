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

// Hexadecimal2Ascii (s string) string
func Hexadecimal2Ascii(s string) string {
	return fmt.Sprintf("%q", ValueOf(s))
}

// Hexadecimal2Binary (s string) string
func Hexadecimal2Binary(s string) string {
	i, err := utils.Parse(s, utils.HEXADECIMAL_BASE)
	utils.CheckError(err, utils.HEXADECIMAL, utils.BINARY)
	s = strconv.Itoa(int(i))
	return decimal.Decimal2Binary(s)
}

// Hexadecimal2Octal (s string) string
func Hexadecimal2Octal(s string) string {
	i, err := utils.Parse(s, utils.HEXADECIMAL_BASE)
	utils.CheckError(err, utils.HEXADECIMAL, utils.OCTAL)
	s = strconv.Itoa(int(i))
	return decimal.Decimal2Octal(s)
}

// Hexadecimal2Decimal (s string) string
func Hexadecimal2Decimal(s string) string {
	i, err := utils.Parse(s, utils.HEXADECIMAL_BASE)
	utils.CheckError(err, utils.HEXADECIMAL, utils.DECIMAL)
	return strconv.Itoa(int(i))
}

// String (s string) string
func String(s string) string {
	return fmt.Sprintf("%x", ValueOf(s))
}

// ValueOf (s string) int
func ValueOf(s string) int {
	i, err := utils.Parse(s, utils.HEXADECIMAL_BASE)
	utils.CheckError(err, utils.STRING, utils.HEXADECIMAL)
	return int(i)
}
