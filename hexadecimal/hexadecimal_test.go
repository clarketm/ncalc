/*

Copyright 2018 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

*/

package hexadecimal_test

import (
	"fmt"

	"github.com/clarketm/ncalc/hexadecimal"
)

func Example() {

	// HEXADECIMAL
	h := "64"
	fmt.Println(hexadecimal.Hexadecimal2Ascii(h))
	fmt.Println(hexadecimal.Hexadecimal2Binary(h))
	fmt.Println(hexadecimal.Hexadecimal2Octal(h))
	fmt.Println(hexadecimal.Hexadecimal2Decimal(h))
	fmt.Printf("%v:%T", hexadecimal.String(h), hexadecimal.String(h))

	// Output:
	// 'd'
	// 1100100
	// 144
	// 100
	// 64:string
}
