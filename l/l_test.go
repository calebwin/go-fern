package l

import (
  "testing"
)

func TestL(t *testing.T) {
  var myL L = generate("ABC", map[string]string{"A": "AB"})
  if iterate(myL, 3) != "ABBBBC" {
    t.Fail()
  }
}
