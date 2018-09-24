# l
A tiny framework for creating Lindenmayer systems. Currently supports stochastic grammars, context-sensitive grammars.

# Usage
```
import "github.com/calebwin/l"

var myL = l.generate("ABC") // create new L-System with an axiom of "ABC"

myRules := map[string][]l.Successor {
  "A" : []l.Successor{
    l.Successor {
      "AB",
      1.0,
      "",
      "",
    },
  },
  "B" : []l.Successor{
    l.Successor {
      "BA",
      1.0,
      "C",
      "A",
    },
  },
}
myL = l.setRules(myL, myRules) // set rules for the L-System

l.iterate(myL, 3) // "ABBBBC" after 3 iterations of the L-System
```
