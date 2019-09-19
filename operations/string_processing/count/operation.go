package count

import (
	"strings"

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

	a.logger.Debug("Executing operation...", input.S0, "   ", input.S1)
	a.logger.Infof("Counting substrings, '%s', in '%s'...", input.S1, input.S0)

	c := strings.Count(input.S0, input.S1)

	return c, nil
}
