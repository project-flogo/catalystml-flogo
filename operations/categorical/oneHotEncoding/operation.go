package oneHotEncoding

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
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

	in.FromMap(inputs)

	params, err := coerce.ToArray(a.params.Columns)
	if err != nil {
		return nil, err
	}
	for _, val := range params {
		arr, err := coerce.ToArray(in.Data[val.(string)])
		if err != nil {
			return nil, err
		}
		temp := removeDuplicate(arr)

		for _, newKey := range temp {
			var result []interface{}

			for _, data := range arr {
				if newKey == data {
					result = append(result, 1)
				} else {
					result = append(result, 0)
				}
			}
			in.Data[newKey.(string)] = result
		}
		delete(in.Data, val.(string))
	}
	a.logger.Debug("Output of hot encoding..", in.Data)
	
	return in.Data, nil
}

func removeDuplicate(arr []interface{}) (result []interface{}) {
	tempMap := make(map[string]interface{})

	for _, val := range arr {

		if _, ok := tempMap[val.(string)]; !ok {
			tempMap[val.(string)] = "1"
		}
	}
	for key, _ := range tempMap {
		result = append(result, key)
	}
	return result
}
