package expreduce

import (
	"github.com/corywalker/mathbigext"
	"math/big"
)

func bigMathFnOneParam(fn func(*big.Float) *big.Float, onlyPos bool) func(*Expression, *EvalState) Ex {
	return (func(this *Expression, es *EvalState) Ex {
		if len(this.Parts) != 2 {
			return this
		}

		flt, ok := this.Parts[1].(*Flt)
		if ok {
			if !onlyPos || flt.Val.Cmp(big.NewFloat(0)) == 1 {
				return &Flt{fn(flt.Val)}
			}
		}
		return this
	})
}

func GetPowerDefinitions() (defs []Definition) {
	defs = append(defs, Definition{
		Name:    "Power",
		Default: "1",
		toString: func(this *Expression, form string, context *String, contextPath *Expression) (bool, string) {
			return ToStringInfixAdvanced(this.Parts[1:], "^", false, "", "", form, context, contextPath)
		},
		legacyEvalFn: func(this *Expression, es *EvalState) Ex {
			if len(this.Parts) != 3 {
				return this
			}

			// TODO: Handle cases like float raised to the float and things raised to
			// zero and 1

			baseInt, baseIsInt := this.Parts[1].(*Integer)
			powerInt, powerIsInt := this.Parts[2].(*Integer)
			baseFlt, baseIsFlt := this.Parts[1].(*Flt)
			powerFlt, powerIsFlt := this.Parts[2].(*Flt)
			// Anything raised to the 1st power is itself
			if powerIsFlt {
				if powerFlt.Val.Cmp(big.NewFloat(1)) == 0 {
					return this.Parts[1]
				}
			} else if powerIsInt {
				if powerInt.Val.Cmp(big.NewInt(1)) == 0 {
					return this.Parts[1]
				}
			}
			// Anything raised to the 0th power is 1, with a small exception
			powerPositivity := -2
			if powerIsFlt {
				powerPositivity = powerFlt.Val.Cmp(big.NewFloat(0))
			} else if powerIsInt {
				powerPositivity = powerInt.Val.Cmp(big.NewInt(0))
			}
			basePositivity := -2
			if baseIsFlt {
				basePositivity = baseFlt.Val.Cmp(big.NewFloat(0))
			} else if baseIsInt {
				basePositivity = baseInt.Val.Cmp(big.NewInt(0))
			}
			if powerPositivity == 0 && (baseIsInt || baseIsFlt) {
				if basePositivity == 0 {
					return &Symbol{"System`Indeterminate"}
				}
				return &Integer{big.NewInt(1)}
			}
			if powerPositivity == 1 && basePositivity == 0 {
				return this.Parts[1]
			}
			if basePositivity == -1 && powerIsFlt {
				return this
			}

			//es.Debugf("Power eval. baseIsInt=%v, powerIsInt=%v", baseIsInt, powerIsInt)
			// Fully integer Power expression
			if baseIsInt && powerIsInt {
				cmpres := powerInt.Val.Cmp(big.NewInt(0))
				//es.Debugf("Cmpres: %v", cmpres)
				if cmpres == 1 {
					res := big.NewInt(0)
					res.Exp(baseInt.Val, powerInt.Val, nil)
					return &Integer{res}
				} else if cmpres == -1 {
					newbase := big.NewInt(0)
					absPower := big.NewInt(0)
					absPower.Abs(powerInt.Val)
					newbase.Exp(baseInt.Val, absPower, nil)
					if newbase.Cmp(big.NewInt(1)) == 0 {
						return &Integer{big.NewInt(1)}
					}
					if newbase.Cmp(big.NewInt(-1)) == 0 {
						return &Integer{big.NewInt(-1)}
					}
					//return NewExpression([]Ex{&Symbol{"System`Power"}, &Integer{newbase}, &Integer{big.NewInt(-1)}})
					return NewExpression([]Ex{&Symbol{"System`Rational"}, &Integer{big.NewInt(1)}, &Integer{newbase}})
				} else {
					return NewExpression([]Ex{&Symbol{"System`Error"}, &String{"Unexpected zero power in Power evaluation."}})
				}
			}

			if baseIsFlt && powerIsInt {
				return &Flt{mathbigext.Pow(baseFlt.Val, big.NewFloat(0).SetInt(powerInt.Val))}
			}
			if baseIsInt && powerIsFlt {
				return &Flt{mathbigext.Pow(big.NewFloat(0).SetInt(baseInt.Val), powerFlt.Val)}
			}
			if baseIsFlt && powerIsFlt {
				return &Flt{mathbigext.Pow(baseFlt.Val, powerFlt.Val)}
			}
			return this
		},
	})
	defs = append(defs, Definition{
		Name: "PowerExpand",
		// This function is not implemented to a satisfiable extent. Do not
		// document it yet.
		OmitDocumentation: true,
		SimpleExamples: []TestInstruction{
			&TestComment{"`PowerExpand` can expand nested log expressions:"},
			&SameTest{"Log[a] + e (Log[b] + d Log[c])", "PowerExpand[Log[a (b c^d)^e]]"},
		},
		Rules: []Rule{
			{"PowerExpand[exp_]", "exp //. {Log[x_ y_]:>Log[x]+Log[y],Log[x_^k_]:>k Log[x]}"},
		},
	})
	defs = append(defs, Definition{Name: "Expand"})
	defs = append(defs, Definition{
		Name: "Log",
		legacyEvalFn: bigMathFnOneParam(mathbigext.Log, true),
	})
	defs = append(defs, Definition{Name: "Sqrt"})
	defs = append(defs, Definition{Name: "I"})
	defs = append(defs, Definition{Name: "PolynomialQ"})
	defs = append(defs, Definition{Name: "Exponent"})
	defs = append(defs, Definition{Name: "Coefficient"})
	defs = append(defs, Definition{Name: "CoefficientList"})
	defs = append(defs, Definition{Name: "PolynomialQuotientRemainder"})
	defs = append(defs, Definition{Name: "PolynomialQuotient"})
	defs = append(defs, Definition{Name: "PolynomialRemainder"})
	defs = append(defs, Definition{Name: "FactorTermsList"})
	defs = append(defs, Definition{Name: "FactorTerms"})
	defs = append(defs, Definition{Name: "ExpreduceFactorConstant"})
	defs = append(defs, Definition{Name: "Variables"})
	defs = append(defs, Definition{Name: "PolynomialGCD"})
	defs = append(defs, Definition{Name: "SquareFreeQ"})
	defs = append(defs, Definition{
		Name:              "PSimplify",
		OmitDocumentation: true,
		ExpreduceSpecific: true,
	})
	defs = append(defs, Definition{Name: "FactorSquareFree"})
	return
}
