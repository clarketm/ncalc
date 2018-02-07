/*

Copyright 2018 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

*/

package octal_test

import (
	"fmt"

	"github.com/clarketm/ncalc/octal"
)

func Example() {

	// OCTAL
	o := "144"
	fmt.Println(octal.Octal2Ascii(o))
	fmt.Println(octal.Octal2Binary(o))
	fmt.Println(octal.Octal2Decimal(o))
	fmt.Println(octal.Octal2Hexadecimal(o))
	fmt.Printf("%v:%T", octal.String(o), octal.String(o))

	// Output:
	// 'd'
	// 1100100
	// 100
	// 64
	// 144:string
}
