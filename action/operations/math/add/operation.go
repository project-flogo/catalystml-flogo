package add

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"

	"github.com/project-flogo/core/data/coerce"
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

	num0, _ := coerce.ToFloat64(input.Num0)
	num1, _ := coerce.ToFloat64(input.Num1)

	a.logger.Info("Starting operation add.")
	a.logger.Debugf("Adding %f by %f", num0, num1)

	out := num0 + num1

	a.logger.Info("Operation add completed")
	a.logger.Debug("Output of Operation add ", out)
	return out, nil
}
