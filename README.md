# l
A tiny framework for creating Lindenmayer systems. Support is included for stochastic grammars, context-sensitive grammars.

# Usage
```
var myL = generate("ABC")

productionRules := map[string]string{
  "A": "AB",
}
myL = setGrammar(myL, productionRules)

iterate(myL, 3) // "ABBBBC"
```
