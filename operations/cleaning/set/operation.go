package set

import (
	"github.com/project-flogo/cml/action/operation"
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

	a.logger.Info("Executing operation...", input.Arr)
	a.logger.Info("Creating set...")

	set := make(map[interface{}]bool)

	out := make([]interface{}, 0, len(input.Arr))
	for _, blah := range input.Arr {
		if set[blah] == false {
			out = append(out, blah)
			set[blah] = true
		}
	}

	return out, nil
}
