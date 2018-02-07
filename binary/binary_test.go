/*

Copyright 2018 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

*/

package binary_test

import (
	"fmt"

	"github.com/clarketm/ncalc/binary"
)

func Example() {

	// BINARY
	b := "1100100"
	fmt.Println(binary.Binary2Ascii(b))
	fmt.Println(binary.Binary2Octal(b))
	fmt.Println(binary.Binary2Decimal(b))
	fmt.Println(binary.Binary2Hexadecimal(b))
	fmt.Printf("%v:%T", binary.String(b), binary.String(b))

	// Output:
	// 'd'
	// 144
	// 100
	// 64
	// 1100100:string
}
