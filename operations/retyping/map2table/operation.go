package map2table

import (
	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/core/data/coerce"
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

	err := in.FromMap(inputs)

	result, err := convertToTable(in.Map, in.ColOrder, a.params.Axis)

	if err != nil {
		return nil, err
	}
	a.logger.Info("Results...", result)

	return result, nil
}

func convertToTable(inputMap map[string]interface{}, order []interface{}, axis int) ([][]interface{}, error) {

	result := make([][]interface{}, len(order))
	if axis == 0 {
		//row
		for index, ord := range order {

			val, _ := coerce.ToArray(inputMap[ord.(string)])
			result[index] = val
		}

	} else {
		for index, ord := range order {
			temp, _ := coerce.ToArray(inputMap[ord.(string)])

			for key, val := range temp {

				if result[key] == nil {
					result[key] = make([]interface{}, len(order))
				}
				result[key][index] = val
			}
		}
	}
	return result, nil
}
