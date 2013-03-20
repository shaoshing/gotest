package assert

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"runtime/debug"
	"strings"
	"testing"
)

var Test *testing.T

func Nil(actual interface{}) {
	NilM(actual, "")
}

func NilM(actual interface{}, msg string, args ...interface{}) {
	if !isNil(actual) {
		error("Assert nil, but failed", nil, actual, msg, args...)
	}
}

func NotNil(actual interface{}) {
	NotNilM(actual, "")
}

func NotNilM(actual interface{}, msg string, args ...interface{}) {
	if isNil(actual) {
		error("Assert not nil, but failed", "not nil", nil, msg, args...)
	}
}

func isNil(value interface{}) bool {
	val := reflect.ValueOf(value)

	if !val.IsValid() {
		return true
	}

	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return val.Pointer() == 0
	}

	return false
}

func True(actual bool) {
	TrueM(actual, "")
}

func TrueM(actual bool, msg string, args ...interface{}) {
	if !actual {
		error("Assert true, but failed", !actual, actual, msg, args...)
	}
}

func NotTrue(actual bool) {
	NotTrueM(actual, "")
}

func NotTrueM(actual bool, msg string, args ...interface{}) {
	if actual {
		error("Assert not true, but failed", !actual, actual, msg, args...)
	}
}

func Equal(exp, actual interface{}) {
	EqualM(exp, actual, "")
}

func EqualM(exp, actual interface{}, msg string, args ...interface{}) {
	if !equal(exp, actual) {
		error("Assert equal, but failed", exp, actual, msg, args...)
	}
}

func NotEqual(exp, actual interface{}) {
	NotEqualM(exp, actual, "")
}

func NotEqualM(exp, actual interface{}, msg string, args ...interface{}) {
	if equal(exp, actual) {
		error("Assert doesn't equal, but failed", exp, actual, msg, args...)
	}
}

func equal(exp, actual interface{}) bool {
	expVal := reflect.ValueOf(exp)
	if !expVal.IsValid() {
		panic("Please use assert.Nil instead.")
	}

	return reflect.DeepEqual(exp, actual)
}

func Contain(expect, content string) {
	ContainM(expect, content, "")
}

func ContainM(expect, content string, msg string, args ...interface{}) {
	if !contain(expect, content) {
		error("Assert contain text, but failed", "contain: "+expect, content, msg, args...)
	}
}

func NotContain(expect, content string) {
	NotContainM(expect, content, "")
}

func NotContainM(expect, content string, msg string, args ...interface{}) {
	if contain(expect, content) {
		error("Assert doesn't contain text, but failed", "contain: "+expect, content, msg, args...)
	}
}

func contain(expect, content string) bool {
	return strings.Contains(content, expect)
}

func Match(expReg, content string) {
	MatchM(expReg, content, "")
}

func MatchM(expReg, content string, msg string, args ...interface{}) {
	if !match(expReg, content) {
		error("Assert match content, but failed", "match: /"+expReg+"/", content, msg, args...)
	}
}

func NotMatch(expReg, content string) {
	NotMatchM(expReg, content, "")
}

func NotMatchM(expReg, content string, msg string, args ...interface{}) {
	if match(expReg, content) {
		error("Assert doesn't match content, but failed", "match: /"+expReg+"/", content, msg, args...)
	}
}

func match(expReg, content string) bool {
	exp := regexp.MustCompile(expReg)
	return exp.Match([]byte(content))
}

func Panic(f func()) {
	PanicM(f, "")
}

func PanicM(f func(), msg string, args ...interface{}) {
	if noPanic(f) {
		error("Assert panic, but didn't", "panic", "no panic", msg, args...)
	}
}

func NoPanic(f func()) {
	NoPanicM(f, "")
}

func NoPanicM(f func(), msg string, args ...interface{}) {
	if !noPanic(f) {
		error("Assert no panic, but panicked", "no panic", "panicked", msg, args...)
	}
}

func noPanic(f func()) bool {
	defer func() {
		recover()
	}()
	f()
	return true
}

func error(message string, exp, actual interface{}, customMsg string, args ...interface{}) {
	if Test == nil {
		println("panic")
		panic("Must assign assert.Test")
	}

	if len(customMsg) != 0 {
		message = fmt.Sprintf(customMsg, args...)
	}

	wd, _ := os.Getwd()
	stack := strings.Replace(strings.Join(strings.Split(string(debug.Stack()), "\n")[4:], "\n"), wd+"/", "./", -1)

	Test.Errorf(`%s
== expect
%s
== actual
%s
== stack
%s
`, message, fmt.Sprint(exp), fmt.Sprint(actual), stack)
	Test.FailNow()
}
