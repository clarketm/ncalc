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

// IsAscii (f string) bool
func IsAscii(f string) bool {
	return f == ASCII || f == string(ASCII[0])
}

// IsNumeric (f string) bool
func IsNumeric(f string) bool {
	return f != ASCII && f != string(ASCII[0])
}

// CheckType (v interface{}, f string)
func CheckType(v interface{}, f string) {

	switch v.(type) {
	case int:
		if IsAscii(f) {
			fmt.Fprintln(os.Stderr, "Invalid ascii character", fmt.Sprintf("%d", v.(int)))
			os.Exit(1)
		}
	case int32:
		if IsNumeric(f) {
			fmt.Fprintln(os.Stderr, "Invalid integer", fmt.Sprintf("%q", v.(int32)))
			os.Exit(1)
		}
	}
}
