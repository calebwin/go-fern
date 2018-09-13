package l

import (
  "testing"
)

func TestL(t *testing.T) {
  productionRules := map[string]string{
    "A": "AB",
  }
  var myL L = generate("ABC", productionRules)
  if iterate(myL, 3) != "ABBBBC" {
    t.Fail()
  }
}
