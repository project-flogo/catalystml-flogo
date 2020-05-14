package contains

import (
	// "strings"

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
	a.logger.Info("Starting Operation Contians.")
	a.logger.Debug("Input for operation contains...", input.S0, " to ", input.S1)

	out := strings.Contains(input.S0, input.S1)
	a.logger.Info("Operation contains completed.")
	a.logger.Debug("Output of Operation contains.", out)

	return out, nil
}
