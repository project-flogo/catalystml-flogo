package levenshteinSimilarity

import (
	"github.com/agext/levenshtein"
	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	// params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	s0 := input.S0
	s1 := input.S1

	a.logger.Info("Executing operation...", s0, s1)

	par := levenshtein.NewParams()
	out := levenshtein.Similarity(s0, s1, par)

	return out, nil
}
