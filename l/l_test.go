package l

import (
  "testing"
)

func TestBasic(t *testing.T) {
  var myL = generate("ABC")

  productionRules := map[string]string{
    "A": "AB",
  }
  myL = setGrammar(myL, productionRules)

  if iterate(myL, 3) != "ABBBBC" {
    t.Fail()
  }
}

// func TestContext(t *testing.T) {
//   productionRules := map[string]string{
//     "A": "AB",
//   }
//   var myL L = generate("ABC", productionRules)
//
//   myL = setContext(myL, map[string]string{"A" : "AB",}, "", "B")
//
//   if iterate(myL, 3) != "ABBBBC" {
//     t.Fail()
//   }
// }
