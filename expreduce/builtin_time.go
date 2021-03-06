package expreduce

import "math/big"
import "time"

func GetTimeDefinitions() (defs []Definition) {
	defs = append(defs, Definition{
		Name: "UnixTime",
		legacyEvalFn: func(this *Expression, es *EvalState) Ex {
			if len(this.Parts) != 1 {
				return this
			}

			return &Integer{big.NewInt(time.Now().UTC().Unix())}
		},
	})
	return
}
