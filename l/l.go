package l

import (
  "strings"
  "math/rand"
)

type Rule struct {
  from string
  to string
  weight float32
}

type L struct {
  axiom string
  rules map[string]Rule
}

func generate(axiom string, grammarRules map[string]string) L {
  var rules map[string]Rule = make(map[string]Rule)

  for from, to := range grammarRules {
    grammarRule := Rule {
      to : to,
      weight : 1.0,
    }
    rules[from] = grammarRule
  }

  return L {
    axiom,
    rules,
  }
}

func setAxiom(l L, newAxiom string) L {
  return L {
    newAxiom,
    l.rules,
  }
}

func setRules(l L, newGrammarRules map[string]string) L {
  var newRules map[string]Rule = make(map[string]Rule)

  for from, to := range newGrammarRules {
    newRule := Rule {
      to : to,
      weight : 1.0,
    }
    newRules[from] = newRule
  }

  return L {
    l.axiom,
    newRules,
  }
}

func setRule(l L, newGrammarRule map[string]string) L {
  var newRules map[string]Rule = make(map[string]Rule)

  for from, to := range newGrammarRule {
    newRule := Rule {
      to : to,
      weight : 1.0,
    }
    newRules[from] = newRule
  }

  for from, to := range l.rules {
      newRules[from] = to
  }

  return L {
    l.axiom,
    newRules,
  }
}

func setWeight(l L, newGrammarRule map[string]string, newWeight float32) L {
  var newRules map[string]Rule = make(map[string]Rule)

  for from, to := range newGrammarRule {
    newRule := Rule {
      to : to,
      weight : newWeight,
    }
    newRules[from] = newRule
  }

  for from, to := range l.rules {
      newRules[from] = to
  }

  return L {
    l.axiom,
    newRules,
  }
}

// TODO set axiom parameters

// TODO set parameters

// TODO set context

// TODO set condition

// TODO set production

func iterate(l L, iterations int) string {
  return iterate_with_rules(l.axiom, l.rules, iterations)
}

func iterate_with_rules(str string, rules map[string]Rule, iterations int) string {
  if  iterations == 0 {
    return str
  }

  var newStr strings.Builder
  for _, char := range str {
    // search for rule for current character
    if _, ok := rules[string(char)]; ok {
      var charRule Rule = rules[string(char)]
      var isCharRuleValid bool = charRule.weight == 1.0 || rand.Float32() <= charRule.weight

      if isCharRuleValid {
        newStr.WriteString(charRule.to)
      }
    } else {
      newStr.WriteString(string(char))
    }
  }

  return iterate_with_rules(newStr.String(), rules, iterations - 1)
}
