package utils

import (
	"fmt"
	"os"
)

// CheckError (err error, formatSrc, formatDst string)
func CheckError(err error, formatSrc, formatDst string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error converting", formatSrc, "to", formatDst, ":", err)
		os.Exit(1)
	}
}
