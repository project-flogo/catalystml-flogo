package date

import (
	"time"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)

	if err != nil {
		return nil, err
	}

	return &Operation{params: p, logger: ctx.Logger()}, nil
}

func (operation *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	operation.logger.Info("Starting operation date.")
	operation.logger.Debug("Input for operation date.", input.Data)

	out, err := time.Parse(operation.params.Format, input.Data)

	if nil != err {
		return nil, err
	}

	operation.logger.Info("Operation date completed.")
	operation.logger.Debug("Output of operation date...", out)
	return out, nil
}
