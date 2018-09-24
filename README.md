# fern
A tiny library for creating Lindenmayer systems. Currently supports stochastic grammars, context-sensitive grammars.

# Usage
```
import "github.com/calebwin/fern"

var myL = fern.generate("ABC") // create new L-System with an axiom of "ABC"

myRules := map[string][]fern.Successor {
  "A" : []fern.Successor{
    fern.Successor {
      "AB",
      1.0,
      "",
      "",
    },
  },
  "B" : []fern.Successor{
    fern.Successor {
      "BA",
      1.0,
      "C",
      "A",
    },
  },
}
myL = fern.setRules(myL, myRules) // set rules for the L-System

fern.iterate(myL, 3) // "ABBBBC" after 3 iterations of the L-System
```
