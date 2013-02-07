package assert

import (
	"reflect"
	"regexp"
	"runtime"
	"testing"
)

var Test *testing.T

func True(actual bool) {
	if !actual {
		error("Assert true, but failed", !actual, actual)
	}
}

func NotTrue(actual bool) {
	if actual {
		error("Assert not true, but failed", !actual, actual)
	}
}

func Equal(exp, actual interface{}) {
	if !equal(exp, actual) {
		error("Assert equal, but failed", exp, actual)
	}
}

func NotEqual(exp, actual interface{}) {
	if equal(exp, actual) {
		error("Assert doesn't equal, but failed", exp, actual)
	}
}

func equal(exp, actual interface{}) bool {
	return reflect.DeepEqual(exp, actual)
}

func Match(expReg, content string) {
	if !match(expReg, content) {
		error("Assert match content, but failed", "match: /"+expReg+"/", content)
	}
}

func NotMatch(expReg, content string) {
	if match(expReg, content) {
		error("Assert doesn't match content, but failed", "match: /"+expReg+"/", content)
	}
}

func match(expReg, content string) bool {
	exp := regexp.MustCompile(expReg)
	return exp.Match([]byte(content))
}

func error(message string, exp, actual interface{}) {
	if Test == nil {
		panic("Must assign assert.Test")
	}

	_, file, line, _ := runtime.Caller(2)
	Test.Errorf("%s\n%s:%d\n== expect ==\n%s\n== actual ==\n%s\n============\n", message, file, line, exp, actual)
	Test.FailNow()
}
