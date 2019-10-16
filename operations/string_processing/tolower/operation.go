package tolower

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

	a.logger.Info("Starting Operation toLower")
	a.logger.Debug("Input for operation tolower.", input.Str)

	out := strings.ToLower(input.Str)

	a.logger.Info("Operation toLower completed.")
	a.logger.Debug("Output for operation tolower.", out)

	return out, nil
}
