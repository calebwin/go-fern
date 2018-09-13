package l

import "strings"

type L struct {
  axiom string
  rules map[string]string
}

func generate(axiom string, rules map[string]string) L {
  return L {
    axiom,
    rules,
  }
}

func iterate(l L, iterations int) string {
  return iterate_rules(l.axiom, l.rules, iterations)
}

func iterate_rules(str string, rules map[string]string, iterations int) string {
  if  iterations == 0 {
    return str
  }

  var newStr strings.Builder
  for _, char := range str {
    if _, ok := rules[string(char)]; ok {
      newStr.WriteString(rules[string(char)])
    } else {
      newStr.WriteString(string(char))
    }
  }

  return iterate_rules(newStr.String(), rules, iterations - 1)
}
