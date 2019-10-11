package split

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

	a.logger.Info("Starting operation split.")
	a.logger.Debug("Output for operation split.")

	out := strings.Split(input.Str, input.Sep)

	a.logger.Info("Operation Split completed.")
	a.logger.Debug("Output of operation split...", out)
	return out, nil
}
