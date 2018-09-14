# l
A tiny framework for creating Lindenmayer systems. Support is included for stochastic grammars, context-sensitive grammars.

# Usage
```
var myL = generate("ABC") // create new L-System with an axiom of "ABC"

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
myL = setRules(myL, myRules) // set rules for the L-System

iterate(myL, 3) // "ABBBBC" after 3 iterations of the L-System
```
