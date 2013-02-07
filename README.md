
## Examples

```go
package assert

import (
  "github.com/shaoshing/gotest"
  "testing"
)

func TestAssertions(t *testing.T) {
  Test = t

  assert.True(true)
  assert.NotTrue(false)

  assert.Equal(true, true)
  assert.Equal(123, 123)
  assert.Equal("123", "123")

  assert.NotEqual(true, false)
  assert.NotEqual(123, 456)
  assert.NotEqual("123", "456")

  assert.Match("123", "123456")
  assert.Match("56$", "123456")
  assert.NotMatch("789", "123456")
  assert.NotMatch("123$", "123456")
}
```
