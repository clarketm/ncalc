/*

Copyright 2017 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

NAME:
    ncalc – number base converter.

SYNOPSIS:
    ncalc [ opts... ] [ number|character ]

OPTIONS:
    -h, --help                  print usage.
    -i, --input format          input format. see FORMATS. (default: decimal|ascii)
    -o, --output format         output format. see FORMATS. (default: all)
    -v, --version               print version number.

FORMATS:
    (a)scii                     character
    (b)inary                    base 2
    (o)ctal                     base 8
    (d)ecimal                   base 10
    (h)exadecimal               base 16

EXAMPLES:
    ncalc "6"                               # output `decimal` number `6` in `all` formats
    ncalc "G"                               # output `ascii` character `G` in `all` formats
    ncalc -i a "f"                          # output `ascii` character `f` in `all` formats
    ncalc -i decimal -o ascii "15"          # output `decimal` number `15` as `ascii`
    ncalc --input h --output o "ff"         # output `hexadecimal` number `ff` as `octal`

*/

package main

import (
	"fmt"
	"os"

	flag "github.com/clarketm/pflag"

	"github.com/clarketm/ncalc/ascii"
	"github.com/clarketm/ncalc/binary"
	"github.com/clarketm/ncalc/decimal"
	"github.com/clarketm/ncalc/hexadecimal"
	"github.com/clarketm/ncalc/octal"
	"github.com/clarketm/ncalc/utils"
	"github.com/fatih/color"
)

// VERSION - current version number
const VERSION = "v1.0.2"

type inputFlag []string

func (i *inputFlag) String() string {
	return "decimal|ascii"
}

func (i *inputFlag) Type() string {
	return "string"
}

func (i *inputFlag) Set(value string) error {
	*i = getFormat(value)
	return nil
}

type outputFlag []string

func (o *outputFlag) String() string {
	return "all"
}
func (o *outputFlag) Type() string {
	return "string"
}

func (o *outputFlag) Set(value string) error {
	*o = getFormat(value)
	return nil
}

// Flags
var version bool

var inputFormat inputFlag
var outputFormat outputFlag = utils.ALL

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
	flag.VarP(&inputFormat, "input", "i", "input `format`: see FORMATS.")

	// -o, --output
	flag.VarP(&outputFormat, "output", "o", "output `format`: see FORMATS.")

	// -v, --version
	flag.BoolVarP(&version, "version", "v", false, "print version number")

	// Usage
	flag.Usage = func() {
		println()
		fmt.Printf("NAME:\n")
		fmt.Printf("\tncalc – number base converter.\n")
		println()
		fmt.Printf("SYNOPSIS:\n")
		fmt.Printf("\t%v [ opts... ] [ number|character ]\n", bold("ncalc"))
		println()
		fmt.Printf("OPTIONS:\n")
		flag.PrintDefaults()
		println()
		fmt.Printf("FORMATS:\n")
		fmt.Printf("\t(a)scii      \tcharacter\n")
		fmt.Printf("\t(b)inary     \tbase 2\n")
		fmt.Printf("\t(o)ctal      \tbase 8\n")
		fmt.Printf("\t(d)ecimal    \tbase 10\n")
		fmt.Printf("\t(h)exadecimal\tbase 16\n")
		println()
		os.Exit(statusCode)
	}
}

func printVersion() {
	fmt.Printf("\n%s %v\n", bold("Version:"), VERSION)
	os.Exit(0)
}

func getFormat(format string) []string {
	o := make([]string, 1)

	switch format {
	case utils.ASCII, string(utils.ASCII[0]):
		o = []string{utils.ASCII}
	case utils.BINARY, string(utils.BINARY[0]):
		o = []string{utils.BINARY}
	case utils.OCTAL, string(utils.OCTAL[0]):
		o = []string{utils.OCTAL}
	case utils.DECIMAL, string(utils.DECIMAL[0]):
		o = []string{utils.DECIMAL}
	case utils.HEXADECIMAL, string(utils.HEXADECIMAL[0]):
		o = []string{utils.HEXADECIMAL}
	default:
		fmt.Fprintln(os.Stderr, "Unkown format", format)
		os.Exit(1)
	}
	return o
}

func setDefaultInputFormat(v interface{}) {
	if utils.IsDecimal(v.(string)) {
		inputFormat = []string{utils.DECIMAL}
	} else {
		inputFormat = []string{utils.ASCII}
	}
}

// main ()
func main() {
	flag.Parse()

	if version {
		printVersion() // version and EXIT
	}

	if len(flag.Args()) < 1 {
		flag.Usage() // usage and EXIT
	}

	arg := flag.Args()[0] // extract arg

	if len(inputFormat) < 1 {
		setDefaultInputFormat(arg)
	}

	// DEBUG
	// fmt.Println("input", inputFormat)
	// fmt.Println("output", outputFormat)
	// fmt.Println("arg", arg)

	println()
	for _, o := range outputFormat {
		fn := funcMap[string(inputFormat[0])+"|"+string(o)]
		result := utils.Invoke(fn, arg)
		fmt.Printf("%v: %v\n", bold(o), result)
	}
	println()

}
