package utils

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
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
		if VERBOSE {
			fmt.Fprintln(os.Stderr, "There was an error converting", formatSrc, "to", formatDst, fmt.Sprintf("\n%v", err))
		} else {
			fmt.Fprintln(os.Stderr, "There was an error converting", formatSrc, "to", formatDst)
		}
		os.Exit(1)
	}
}

// IsAscii (f string) bool
func IsAscii(s string) bool {
	matched, _ := regexp.MatchString("[[:alpha:]]", s)
	return matched
}

// IsNumeric (f string) bool
func IsNumeric(s string) bool {
	var err error
	if _, err = strconv.ParseFloat(s, 0); err == nil {
		return true
	}
	if _, err = strconv.ParseInt(s, 0, 0); err == nil {
		return true
	}
	if _, err = strconv.ParseInt(s, BINARY_BASE, 0); err == nil {
		return true
	}
	if _, err = strconv.ParseInt(s, OCTAL_BASE, 0); err == nil {
		return true
	}
	if _, err = strconv.ParseInt(s, DECIMAL_BASE, 0); err == nil {
		return true
	}
	if _, err = strconv.ParseInt(s, HEXADECIMAL_BASE, 0); err == nil {
		return true
	}
	return false
}

// IsDecimal (f string) bool
func IsDecimal(s string) bool {
	if _, err := strconv.ParseInt(s, DECIMAL_BASE, 0); err == nil {
		return true
	}
	return false
}

// IsLiteral (s string) bool
func IsLiteral(s string) bool {
	var octalLiteral = regexp.MustCompile(`^0[0-7]+$`)
	var hexLiteral = regexp.MustCompile(`^0[xX][0-9a-zA-Z]+$`)
	return octalLiteral.MatchString(s) || hexLiteral.MatchString(s)
}

// Parse (s string, base int) (int64, error)
func Parse(s string, base int) (int64, error) {
	if IsLiteral(s) {
		return strconv.ParseInt(s, 0, 0)
	} else {
		return strconv.ParseInt(s, base, 0)
	}
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
