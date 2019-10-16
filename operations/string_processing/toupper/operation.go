package toupper

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

	a.logger.Info("Starting Operation toupper.")
	a.logger.Debug("Input for operation toupper.", input.Str)

	out := strings.ToUpper(input.Str)

	a.logger.Info("Operation toupper completed.")
	a.logger.Debug("Output of operation toupper...", out)

	return out, nil
}
