package l

import (
  "strings"
  "math/rand"
  // "fmt"
)

type Successor struct {
  successor string
  weight float32
  leftContext string
  rightContext string
  // TODO successorFn
  // TODO conditionFn
}

type L struct {
  axiom string
  rules map[string][]Successor
}

func generate(axiom string) L {
  return L {
    axiom,
    make(map[string][]Successor),
  }
}

func setAxiom(currL L, newAxiom string) L {
  return L {
    newAxiom,
    currL.rules,
  }
}

func setRules(currL L, newRules map[string][]Successor) L {
  currLNewRules := make(map[string][]Successor)

  for from, currSuccessors := range currL.rules {
    currLNewRules[from] = currSuccessors
  }

  for from, newSuccessors := range newRules {
    currLNewRules[from] = newSuccessors
  }

  return L {
    currL.axiom,
    currLNewRules,
  }
}

func iterate(currL L, numIterations int) string {
  return rewrite(currL.axiom, currL.rules, numIterations)
}

func rewrite(str string, rules map[string][]Successor, numIterations int) string {
  if numIterations == 0 {
    return str
  }

  var newStr strings.Builder

  for _, char := range str {
    var charStr string = string(char)
    var charIndex int = strings.Index(str, charStr)

    // check if charStr has a valid successors
    charStrSuccessors := rules[charStr]
    if len(charStrSuccessors) > 0 {
      var successorIndex = -1

      // calculate total weight of all successors, find deterministic successor
      var totalWeight float32 = 0.0
      for i, charStrSuccessor := range charStrSuccessors {
        totalWeight += charStrSuccessor.weight
        if charStrSuccessor.weight == 1.0 {
          successorIndex = i
        }
      }

      // random successor
      if successorIndex == -1 {
        var randWeight float32 = rand.Float32()
        var cumulativeWeight float32 = 0.0
        for i, charStrSuccessor := range charStrSuccessors {
          if randWeight < cumulativeWeight / totalWeight {
            successorIndex = i
          }

          cumulativeWeight += charStrSuccessor.weight
        }
      }

      validSuccessor := charStrSuccessors[successorIndex]

      // ensure context is correct
      if validSuccessor.leftContext != "" && (charIndex - len(validSuccessor.leftContext) < 0 || str[charIndex - len(validSuccessor.leftContext) : charIndex] != validSuccessor.leftContext) {
        successorIndex = -1
      }
      if validSuccessor.rightContext != "" && (charIndex + len(validSuccessor.rightContext) >= len(str) || str[charIndex + 1 : charIndex + len(validSuccessor.rightContext) + 1] != validSuccessor.rightContext) {
        successorIndex = -1
      }

      if successorIndex > -1 {
        newStr.WriteString(charStrSuccessors[0].successor)
      } else {
        newStr.WriteString(charStr)
      }
    } else {
      newStr.WriteString(charStr)
    }
  }

  return rewrite(newStr.String(), rules, numIterations - 1)
}

// TODO final
