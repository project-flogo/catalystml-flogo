package trimright

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

	a.logger.Info("Starting Operation TrimRight.")
	a.logger.Debug("Input for operation TrimRight.", input.S0, " ", input.S1)

	out := strings.TrimRight(input.S0, input.S1)

	a.logger.Info("Operation TrimRight completed.")
	a.logger.Debug("Output for operation TrimRight.", out)

	return out, nil
}
