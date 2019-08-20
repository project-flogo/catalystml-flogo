package dropCol

import (
	"github.com/project-flogo/cml/action/operation"
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

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	in := &Input{}

	in.FromMap(inputs)

	a.logger.Debug("Input of DropCol is ....", in.Data, a.params)
	for _, val := range a.params.Columns {

		if _, ok := in.Data[val.(string)]; ok {
			delete(in.Data, val.(string))
		}

	}
	a.logger.Debug("Output of DropCol is ....", in.Data)
	return in.Data, nil
}
