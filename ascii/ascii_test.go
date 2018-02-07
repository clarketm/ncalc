/*

Copyright 2018 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

*/

package ascii_test

import (
	"fmt"

	"github.com/clarketm/ncalc/ascii"
)

func Example() {
	// ASCII
	a := "d"
	fmt.Println(ascii.Ascii2Binary(a))
	fmt.Println(ascii.Ascii2Octal(a))
	fmt.Println(ascii.Ascii2Decimal(a))
	fmt.Println(ascii.Ascii2Hexadecimal(a))
	fmt.Printf("%v:%T", ascii.String(a), ascii.String(a))

	// Output:
	// 1100100
	// 144
	// 100
	// 64
	// 'd':string
}
