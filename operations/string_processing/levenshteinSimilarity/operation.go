package levenshteinSimilarity

import (
	"github.com/agext/levenshtein"
	"github.com/project-flogo/catalystml-flogo/action/operation"
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

	a.logger.Info("Starting operation Levenshtein Similarity.")
	a.logger.Debug("Input for operation levenshtein similarity.", s0, s1)

	par := levenshtein.NewParams()
	out := levenshtein.Similarity(s0, s1, par)

	a.logger.Info("Operation Levenshtein Similarity completed.")
	a.logger.Debug("Output for operation levenshtein similarity.", out)

	return out, nil
}
