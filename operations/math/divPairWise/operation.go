package divPairWise

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"

	"github.com/project-flogo/catalystml-flogo/operations/common"

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

	a.logger.Info("Starting operation divPairWise.")

	out, err := common.MatrixPairWiseOperation(input.Matrix0, input.Matrix1, common.MatirxPairWiseMulti)
	if err != nil {
		return nil, err
	}
	a.logger.Info("Operation divPairWise completed")
	a.logger.Debug("Output of Operation divPairWise ", out)
	return out, nil
}
