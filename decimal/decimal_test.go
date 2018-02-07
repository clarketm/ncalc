/*

Copyright 2018 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

*/

package decimal_test

import (
	"fmt"

	"github.com/clarketm/ncalc/decimal"
)

func Example() {

	// DECIMAL
	d := "100"
	fmt.Println(decimal.Decimal2Ascii(d))
	fmt.Println(decimal.Decimal2Binary(d))
	fmt.Println(decimal.Decimal2Octal(d))
	fmt.Println(decimal.Decimal2Hexadecimal(d))
	fmt.Printf("%v:%T", decimal.String(d), decimal.String(d))

	// Output:
	// 'd'
	// 1100100
	// 144
	// 64
	// 100:string
}
