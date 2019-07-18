package operation

import (
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/fps/operation"
)

func init() {
	_ = operation.Register(&Operation{}, New)
}

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

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	a.logger.Info("Executing operation...", input.InputSample, a.params.Sample)
	out := make(map[string]interface{})

	out["sample"] = 26

	out["OutputArray"] = []interface{}{1, 2, 3, 4, 5}

	return out, nil
}