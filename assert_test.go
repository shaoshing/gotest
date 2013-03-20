package assert_test

import (
	"github.com/shaoshing/gotest"
	"testing"
)

func TestAssertions(t *testing.T) {
	assert.Test = t

	assert.Nil(nil)
	assert.NilM(nil, "Each assertion method has an alternative method eneded with M to support custom error message.")
	assert.NilM(nil, "Custome Error Message works like printf: %s", "actual value")

	var nilSlice []int
	var nilPointer *[]int
	assert.NilM(nilSlice, "Nil Slice == nil")
	assert.NilM(nilPointer, "Nil Pointer == nil")

	assert.NotNil(1)
	assert.NotNil("1")
	assert.NotNil(true)

	assert.True(true)
	assert.NotTrue(false)

	assert.Equal(true, true)
	assert.Equal(123, 123)
	assert.Equal("123", "123")

	assert.NotEqual(true, false)
	assert.NotEqual(123, 456)
	assert.NotEqual("123", "456")

	assert.Contain("123", "123456")
	assert.Contain("56", "123456")

	assert.NotContain("789", "123456")
	assert.NotContain("123$", "123456")

	assert.Match("123", "123456")
	assert.Match("56$", "123456")

	assert.NotMatch("789", "123456")
	assert.NotMatch("123$", "123456")

	assert.Panic(func() {
		panic("Test")
	})

	assert.NoPanic(func() {
		// No panic
	})
}

// func TestTrue(t *testing.T) {
// 	Test = t
// 	True(false)
// }

// func TestFailEqual(t *testing.T) {
// 	Test = t
// 	Equal(false, true)
// }

// func TestFailContain(t *testing.T) {
// 	Test = t
// 	Contain("123", "456")
// }

// func TestFailMatch(t *testing.T) {
// 	Test = t
// 	MatchM("123", "456", "Custom message %s", "with params")
// }

// func TestNilAssetion(t *testing.T) {
// 	Test = t
// 	Equal(nil, nil)
// }
