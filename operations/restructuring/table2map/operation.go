package table2map

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

	err := in.FromMap(inputs)
	a.logger.Info("Starting Operation Table to Map.")
	a.logger.Debug("The inputs of Operation Table to Map.", inputs)

	result, err := convertToMap(in.Table, in.ColKeys, a.params.Axis)

	if err != nil {
		return nil, err
	}
	a.logger.Info("Operation Table to Map completed.")
	a.logger.Debug("The output of Operation Table to Map.", result)
	return result, nil
}

func convertToMap(arr []interface{}, keys []interface{}, axis int) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	if axis == 1 {
		for key, val := range arr {

			mapKey, err := coerce.ToString(keys[key])
			if err != nil {
				return nil, err
			}
			result[mapKey] = val

		}
	} else {
		for key, val := range keys {
			mapKey, err := coerce.ToString(val)
			if err != nil {
				return nil, err
			}
			result[mapKey] = getColumVal(0, key, arr)
		}
	}

	return result, nil
}

func getColumVal(row, key int, arr []interface{}) interface{} {
	var result []interface{}

	_, err := coerce.ToFloat32(arr[0])

	if err != nil {

		for i := row; i < len(arr); i++ {
			temp, _ := coerce.ToArray(arr[i])

			result = append(result, getColumVal(i, key, temp))
		}

		return result
	}

	return arr[key]

}
