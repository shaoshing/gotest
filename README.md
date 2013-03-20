# GoTest

[![Build Status](https://travis-ci.org/shaoshing/gotest.png?branch=master)](https://travis-ci.org/shaoshing/gotest)


## Examples

```go
import (
  "github.com/shaoshing/gotest"
  "testing"
)

func TestAssertions(t *testing.T) {
  Test = t

  assert.True(true)
  assert.NilM(nil, "Each assertion method has an alternative method eneded with M to support custom error message.")
  assert.NilM(nil, "Custome Error Message works like printf: %s", "actual value")

  // List of Assertions:
  //
  // assert.Nil
  // assert.NotNil
  // assert.True
  // assert.NotTrue
  // assert.Equal
  // assert.NotEqual
  // assert.Contain
  // assert.NotContain
  // assert.Match
  // assert.NotMatch
  // assert.Panic
  // assert.NoPanic
  //
  // Checkout assert_test.go for more examples
}
```

## License

GoTest is released under the [MIT License](http://www.opensource.org/licenses/MIT).
