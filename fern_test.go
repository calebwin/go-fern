package fern

import (
  "testing"
)

func TestBasic(t *testing.T) {
  var myL = generate("ABC")

  myRules := map[string][]Successor {
    "A" : []Successor{
      Successor {
        "AB",
        1.0,
        "",
        "",
      },
    },
    "B" : []Successor{
      Successor {
        "BA",
        1.0,
        "C",
        "A",
      },
    },
  }

  myL = setRules(myL, myRules)

  if iterate(myL, 1) != "ABBC" {
    t.Fail()
  }

  if iterate(myL, 3) != "ABBBBC" {
    t.Fail()
  }
}

func TestContext(t *testing.T) {
  var myL = generate("ABC")

  myRules := map[string][]Successor {
    "A" : []Successor{
      Successor{
        "AB",
        1.0,
        "",
        "B",
      },
    },
  }

  myL = setRules(myL, myRules)

  if iterate(myL, 1) != "ABBC" {
    t.Fail()
  }

  if iterate(myL, 3) != "ABBBBC" {
    t.Fail()
  }
}
