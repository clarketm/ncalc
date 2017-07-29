/*

Copyright 2017 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

TODO: Manpage

*/

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/clarketm/ncalc/ascii"
	"github.com/clarketm/ncalc/binary"
	"github.com/clarketm/ncalc/decimal"
	"github.com/clarketm/ncalc/hexadecimal"
	"github.com/clarketm/ncalc/octal"
	"github.com/clarketm/ncalc/utils"
	"github.com/fatih/color"
)

// VERSION - current version number
const VERSION = "v1.0.0"

// versionFlag bool
type versionFlag bool

func (v *versionFlag) IsBoolFlag() bool {
	return true
}

func (v *versionFlag) String() string {
	return fmt.Sprint(*v)
}

func (v *versionFlag) Set(value string) error {
	fmt.Printf("\n%s %v\n", bold("Version:"), *v)
	os.Exit(0)
	return nil
}

type inputFlag []string

func (i *inputFlag) String() string {
	return fmt.Sprint(*i)
}

func (i *inputFlag) Set(value string) error {
	*i = getFormat(value)
	return nil
}

type outputFlag []string

func (o *outputFlag) String() string {
	return fmt.Sprint(*o)
}

func (o *outputFlag) Set(value string) error {
	*o = getFormat(value)
	return nil
}

type allFormats []string

// Flags
var version versionFlag
var inputFormat inputFlag
var outputFormat outputFlag = []string{
	utils.ASCII,
	utils.BINARY,
	utils.OCTAL,
	utils.DECIMAL,
	utils.HEXADECIMAL,
}

// Globals
var statusCode int
var bold = color.New(color.Bold).SprintFunc()
var funcMap = map[string]interface{}{
	"ascii|ascii":             ascii.String,
	"ascii|binary":            ascii.Ascii2Binary,
	"ascii|octal":             ascii.Ascii2Octal,
	"ascii|decimal":           ascii.Ascii2Decimal,
	"ascii|hexadecimal":       ascii.Ascii2Hexadecimal,
	"binary|ascii":            binary.Binary2Ascii,
	"binary|binary":           binary.String,
	"binary|octal":            binary.Binary2Octal,
	"binary|decimal":          binary.Binary2Decimal,
	"binary|hexadecimal":      binary.Binary2Hexadecimal,
	"octal|ascii":             octal.Octal2Ascii,
	"octal|binary":            octal.Octal2Binary,
	"octal|octal":             octal.String,
	"octal|decimal":           octal.Octal2Decimal,
	"octal|hexadecimal":       octal.Octal2Hexadecimal,
	"decimal|ascii":           decimal.Decimal2Ascii,
	"decimal|binary":          decimal.Decimal2Binary,
	"decimal|octal":           decimal.Decimal2Octal,
	"decimal|decimal":         decimal.String,
	"decimal|hexadecimal":     decimal.Decimal2Hexadecimal,
	"hexadecimal|ascii":       hexadecimal.Hexadecimal2Ascii,
	"hexadecimal|binary":      hexadecimal.Hexadecimal2Binary,
	"hexadecimal|octal":       hexadecimal.Hexadecimal2Octal,
	"hexadecimal|decimal":     hexadecimal.Hexadecimal2Decimal,
	"hexadecimal|hexadecimal": hexadecimal.String,
}

// init () - initialize command-line flags
func init() {
	// -i, --input
	flag.Var(&inputFormat, "i", "input `format`: (a)scii, (b)inary, (o)ctal, (d)ecimal, (h)exadecimal")

	// -o, --output
	flag.Var(&outputFormat, "o", "output `format`: (a)scii, (b)inary, (o)ctal, (d)ecimal, (h)exadecimal")

	// -v, --version
	flag.Var(&version, "v", "")
	flag.Var(&version, "version", "print version number")

	// Usage
	flag.Usage = func() {
		println()
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		println()
		os.Exit(statusCode)
	}
}

// main ()
func main() {
	flag.Parse()

	// DEBUG
	// fmt.Println("input", inputFormat)
	// fmt.Println("output", outputFormat)

	if len(flag.Args()) < 1 {
		flag.Usage() // usage
	}

	arg := parseArg(flag.Args()[0])

	// DEBUG
	// fmt.Println("arg", arg)

	println()
	for _, o := range outputFormat {
		fn := funcMap[string(inputFormat[0])+"|"+string(o)]
		result := utils.Invoke(fn, arg)
		fmt.Printf("%v: %v\n", bold(o), result)
	}
	println()

}

func parseArg(a string) interface{} {
	i, err := strconv.Atoi(a)
	if err != nil {
		c := rune(a[0])
		return c
	}
	return i
}

func getFormat(format string) []string {
	o := make([]string, 1)

	switch format {
	case utils.ASCII:
	case string(utils.ASCII[0]):
		o = []string{utils.ASCII}
	case utils.BINARY:
	case string(utils.BINARY[0]):
		o = []string{utils.BINARY}
	case utils.OCTAL:
	case string(utils.OCTAL[0]):
		o = []string{utils.OCTAL}
	case utils.DECIMAL:
	case string(utils.DECIMAL[0]):
		o = []string{utils.DECIMAL}
	case utils.HEXADECIMAL:
	case string(utils.HEXADECIMAL[0]):
		o = []string{utils.HEXADECIMAL}
	case utils.ALL:
		o = []string{
			utils.ASCII,
			utils.BINARY,
			utils.OCTAL,
			utils.DECIMAL,
			utils.HEXADECIMAL,
		}
	default:
		fmt.Fprintln(os.Stderr, "Unkown format", format)
		os.Exit(1)
	}
	return o
}
