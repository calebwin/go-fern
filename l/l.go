package l

import (
  "strings"
  "math/rand"
)

type Rule struct {
  to string
  weight float32
  leftContext string
  rightContext string
}

type L struct {
  axiom string
  rules map[string]Rule
}

func generate(axiom string) L {
  return L {
    axiom,
    make(map[string]Rule),
  }
}

func setAxiom(currL L, newAxiom string) L {
  return L {
    newAxiom,
    currL.rules,
  }
}

func setGrammar(currL L, newGrammarRules map[string]string) L {
  var newRules map[string]Rule = make(map[string]Rule)

  for from, rule := range currL.rules {
    newRules[from] = rule
  }

  for from, to := range newGrammarRules {
    grammarRule := currL.rules[from]
    grammarRule.to = to
    grammarRule.weight = 1.0
    newRules[from] = grammarRule
  }

  return L {
    currL.axiom,
    newRules,
  }
}

func setWeights(currL L, newWeightRules map[string]float32) L {
  var newRules map[string]Rule = make(map[string]Rule)

  for from, rule := range currL.rules {
    newRules[from] = rule
  }

  for from, weight := range newWeightRules {
    grammarRule := currL.rules[from]
    grammarRule.weight = weight
    newRules[from] = grammarRule
  }

  return L {
    currL.axiom,
    newRules,
  }
}
func setContexts(currL L, newLeftContextRules map[string]string, newRightContextRules map[string]string) L {
  var newRules map[string]Rule = make(map[string]Rule)

  for from, rule := range currL.rules {
    newRules[from] = rule
  }

  for from, leftContext := range newLeftContextRules {
    grammarRule := currL.rules[from]
    grammarRule.leftContext = leftContext
    newRules[from] = grammarRule
  }

  for from, rightContext := range newRightContextRules {
    grammarRule := currL.rules[from]
    grammarRule.rightContext = rightContext
    newRules[from] = grammarRule
  }

  return L {
    currL.axiom,
    newRules,
  }
}

// TODO set axiom parameters

// TODO set parameters

// TODO set condition

// TODO set production

// TODO set final

func iterate(l L, iterations int) string {
  return iterate_with_rules(l.axiom, l.rules, iterations)
}

func iterate_with_rules(str string, rules map[string]Rule, iterations int) string {
  if  iterations == 0 {
    return str
  }

  var newStr strings.Builder
  for _, char := range str {
    var char_str string = string(char)

    // search for rule for current character
    if _, ok := rules[string(char)]; ok {
      var charIndex int = strings.Index(str, char_str)
      var charRule Rule = rules[char_str]

      var isCharRuleValid bool = charRule.weight == 1.0 || rand.Float32() <= charRule.weight
      if charRule.leftContext != "" && (charIndex - len(charRule.leftContext) < 0 || str[charIndex - len(charRule.leftContext) : charIndex] != charRule.leftContext) {
        isCharRuleValid = false
      }
      if charRule.rightContext != "" && (charIndex + len(charRule.rightContext) >= len(str) || str[charIndex + 1 : charIndex + len(charRule.rightContext) + 1] != charRule.rightContext) {
        isCharRuleValid = false
      }

      if isCharRuleValid {
        newStr.WriteString(charRule.to)
      }
    } else {
      newStr.WriteString(char_str)
    }
  }

  return iterate_with_rules(newStr.String(), rules, iterations - 1)
}
