package assert

import (
	"testing"
)

func TestAssertions(t *testing.T) {
	Test = t

	True(true)
	NotTrue(false)

	Equal(true, true)
	Equal(123, 123)
	Equal("123", "123")

	NotEqual(true, false)
	NotEqual(123, 456)
	NotEqual("123", "456")

	Contain("123", "123456")
	Contain("56", "123456")
	NotContain("789", "123456")
	NotContain("123$", "123456")

	Match("123", "123456")
	Match("56$", "123456")
	NotMatch("789", "123456")
	NotMatch("123$", "123456")
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
// 	Match("123", "456")
// }
