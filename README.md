# GoTest

Gotest provides some useful assertion methods for writing Go tests.

[![Build Status](https://travis-ci.org/shaoshing/gotest.png?branch=master)](https://travis-ci.org/shaoshing/gotest)

## Quick Example

```go
import (
  "github.com/shaoshing/gotest"
  "testing"
)

func TestAssertions(test *testing.T) {
  assert.Test = test  // Assign the test pointer to Gotest which will be used by the assertion methods.

  assert.True(1 == 1)
  assert.TrueM(1 == 1, "1 should equal to 1")  // Each assertion method has an alternative method
                                               // eneded with M to support custom error message.

  assert.Nil(nil)
  assert.Panic(func(){
    panic("I am panicking!")
  })
}
```

## List of Assertions

* assert.Nil
* assert.NotNil
* assert.True
* assert.NotTrue
* assert.Equal
* assert.NotEqual
* assert.Contain
* assert.NotContain
* assert.Match
* assert.NotMatch
* assert.Panic
* assert.NoPanic

Docs: [http://godoc.org/github.com/shaoshing/gotest](http://godoc.org/github.com/shaoshing/gotest)

## License

GoTest is released under the [MIT License](http://www.opensource.org/licenses/MIT).
