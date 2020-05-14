package concatMap

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
	in := &Input{}
	in.FromMap(inputs)

	a.logger.Info("Starting Concat Map operation.")
	a.logger.Debug("Inputs of Concat Map operation", inputs)

	result := make(map[string]interface{})

	for i := 0; i < len(in.Data); i++ {

		if in.Data[i] != nil {

			for key, val := range in.Data[i].(map[string]interface{}) {
				if result[key] == nil {
					result[key] = make([]interface{},0)
				}
				result[key] = append(result[key].([]interface{}), val)
			}
		}

	}

	a.logger.Info("Operation Concat Map Completed.")
	a.logger.Debug("Output of operation Concat Map.", result)

	return result, nil
}
