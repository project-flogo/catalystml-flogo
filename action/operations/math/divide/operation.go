package divide

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

	num, _ := coerce.ToFloat64(input.Num)
	denom, _ := coerce.ToFloat64(input.Denom)

	a.logger.Info("Starting operation divide.")
	a.logger.Debugf("Dividing %f by %f", num, denom)

	out := num / denom

	a.logger.Info("Operation divide completed")
	a.logger.Debug("Output of Operation divide ", out)
	return out, nil
}
