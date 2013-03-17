package assert

import (
	"testing"
)

func TestAssertions(t *testing.T) {
	Test = t

	True(true)
	TrueM(true, "Should be")

	NotTrue(false)
	NotTrueM(false, "false")

	Equal(true, true)
	Equal(123, 123)
	Equal("123", "123")
	EqualM("123", "123", "123 should equals to 123, %s", "works like fmt.Sprintf")

	NotEqual(true, false)
	NotEqual(123, 456)
	NotEqual("123", "456")
	NotEqualM("123", "456", "123 should not equal to 456")

	Contain("123", "123456")
	Contain("56", "123456")
	ContainM("56", "123456", "123456 should contain 56")

	NotContain("789", "123456")
	NotContain("123$", "123456")
	NotContainM("123$", "123456", "123456 should not contain 123$")

	Match("123", "123456")
	Match("56$", "123456")
	MatchM("56$", "123456", "Reg match")

	NotMatch("789", "123456")
	NotMatch("123$", "123456")
	NotMatchM("123$", "123456", "Reg match")
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
