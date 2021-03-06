package expreduce

import (
	"math/big"
)

// Used for the IntegerPartitions builtin
func genIntegerPartitions(n int, k int, startAt int, prefix []int, parts *[][]int) {
	if len(prefix)+1 > k {
		return
	}
	prefix = append(prefix, 0)
	for i := startAt; i > 0; i-- {
		prefix[len(prefix)-1] = i
		if i == n {
			*parts = append(*parts, make([]int, len(prefix)))
			copy((*parts)[len(*parts)-1], prefix)
		} else {
			genIntegerPartitions(n-i, k, Min(i, n-i), prefix, parts)
		}
	}
}

// Used for the Permutations builtin
func permListContains(permList [][]Ex, perm []Ex, cl *CASLogger) bool {
	for _, permInList := range permList {
		if len(permInList) != len(perm) {
			continue
		}
		same := true
		for i := range perm {
			if !IsSameQ(perm[i], permInList[i], cl) {
				same = false
				//break
			}
		}
		if same {
			return true
		}
	}
	return false
}

// Used for the Permutations builtin
func genPermutations(parts []Ex, cl *CASLogger) (perms [][]Ex) {
	// Base case
	if len(parts) == 1 {
		return [][]Ex{parts}
	}
	// Recursion
	toReturn := [][]Ex{}
	for i, first := range parts {
		// We must make a copy of "parts" because selecting "others" actually
		// modifies "parts" and corrupts it.
		copyParts := make([]Ex, len(parts))
		copy(copyParts, parts)
		others := append(copyParts[:i], copyParts[i+1:]...)
		// TODO: This might be bad for memory complexity.
		otherPerms := genPermutations(others, cl)
		for _, perm := range otherPerms {
			prepended := make([]Ex, len(perm))
			copy(prepended, perm)
			perm = append([]Ex{first}, perm...)
			// TODO: And this is bad for time complexity:
			if !permListContains(toReturn, perm, cl) {
				toReturn = append(toReturn, perm)
			}
		}
	}
	return toReturn
}

// Used for the Factorial builtin
func factorial(n *big.Int) (result *big.Int) {
	result = new(big.Int)

	switch n.Cmp(&big.Int{}) {
	case -1, 0:
		result.SetInt64(1)
	default:
		result.Set(n)
		var one big.Int
		one.SetInt64(1)
		result.Mul(result, factorial(n.Sub(n, &one)))
	}
	return
}

func getCombinatoricsDefinitions() (defs []Definition) {
	defs = append(defs, Definition{
		Name: "IntegerPartitions",
		legacyEvalFn: func(this *Expression, es *EvalState) Ex {
			if len(this.Parts) != 2 && len(this.Parts) != 3 {
				return this
			}

			n, nIsInt := this.Parts[1].(*Integer)
			if !nIsInt {
				return this
			}
			nMachine := int(n.Val.Int64())

			kMachine := nMachine
			if len(this.Parts) == 3 {
				k, kIsInt := this.Parts[2].(*Integer)
				if !kIsInt {
					return this
				}
				kMachine = int(k.Val.Int64())
			}

			cmpVal := n.Val.Cmp(big.NewInt(0))
			if cmpVal == -1 {
				return NewExpression([]Ex{&Symbol{"System`List"}})
			} else if cmpVal == 0 {
				return NewExpression([]Ex{&Symbol{"System`List"}, NewExpression([]Ex{&Symbol{"System`List"}})})
			}

			var parts [][]int
			genIntegerPartitions(nMachine, kMachine, nMachine, []int{}, &parts)

			exParts := NewExpression([]Ex{&Symbol{"System`List"}})
			for _, partition := range parts {
				toAppend := NewExpression([]Ex{&Symbol{"System`List"}})
				for _, integer := range partition {
					toAppend.Parts = append(toAppend.Parts, &Integer{big.NewInt(int64(integer))})
				}
				exParts.Parts = append(exParts.Parts, toAppend)
			}

			return exParts
		},
	})
	defs = append(defs, Definition{
		Name: "Permutations",
		legacyEvalFn: func(this *Expression, es *EvalState) Ex {
			if len(this.Parts) != 2 {
				return this
			}

			list, listIsExpr := this.Parts[1].(*Expression)
			if !listIsExpr {
				return this
			}

			perms := genPermutations(list.Parts[1:], &es.CASLogger)

			exPerms := NewExpression([]Ex{&Symbol{"System`List"}})
			for _, perm := range perms {
				toAppend := NewExpression([]Ex{&Symbol{"System`List"}})
				for _, ex := range perm {
					toAppend.Parts = append(toAppend.Parts, ex)
				}
				exPerms.Parts = append(exPerms.Parts, toAppend)
			}

			return exPerms
		},
	})
	defs = append(defs, Definition{
		Name: "Multinomial",
	})
	defs = append(defs, Definition{
		Name: "Factorial",
		legacyEvalFn: func(this *Expression, es *EvalState) Ex {
			if len(this.Parts) != 2 {
				return this
			}
			asInt, isInt := this.Parts[1].(*Integer)
			if isInt {
				if asInt.Val.Cmp(big.NewInt(0)) == -1 {
					return &Symbol{"System`ComplexInfinity"}
				}
				return &Integer{factorial(asInt.Val)}
			}
			return this
		},
	})
	return
}
