package utils

import (
	"fmt"
	"os"
	"reflect"
)

// Invoke (fn interface{}, args ...interface{}) interface{}
func Invoke(fn interface{}, args ...interface{}) interface{} {
	v := reflect.ValueOf(fn)
	rargs := make([]reflect.Value, len(args))
	for i, a := range args {
		rargs[i] = reflect.ValueOf(a)
	}
	defer recover()
	return v.Call(rargs)[0]
}

// CheckError (err error, formatSrc, formatDst string)
func CheckError(err error, formatSrc, formatDst string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error converting", formatSrc, "to", formatDst, ":", err)
		os.Exit(1)
	}
}
