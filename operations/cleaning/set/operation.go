package set

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {

	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	a.logger.Info("Starting Operation Set.")
	a.logger.Info("Input of Operation Set.", inputs)

	set := make(map[interface{}]bool)

	out := make([]interface{}, 0, len(input.Arr))
	for _, blah := range input.Arr {
		if set[blah] == false {
			out = append(out, blah)
			set[blah] = true
		}
	}
	a.logger.Info("Operation Set Completed")
	a.logger.Debug("Output of Operation Set.", out)

	return out, nil
}
