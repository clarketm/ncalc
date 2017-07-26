/*

Copyright 2017 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

TODO: Manpage

*/

package main

import (
	"fmt"
	"os"
)

// main ()
func main() {

}

// checkError (err error, formatSrc, formatDst string)
func checkError(err error, formatSrc, formatDst string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error converting", formatSrc, "to", formatDst, ":", err)
		os.Exit(1)
	}
}
