package map2table

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	params *Params
	logger log.Logger
}

var ValLen int
var givenType data.Type

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
	a.logger.Info("Starting Operation Map to Table.")
	a.logger.Debug("The inputs of Operation Map to Table.", inputs)
	for _, ord := range in.ColOrder {
		if val, ok := in.Map[ord.(string)]; ok {
			ValLen = len(val.([]interface{}))
			givenType, _ = data.GetType(val.([]interface{})[0])
			break
		}
	}

	result, err := convertToTable(in.Map, in.ColOrder, a.params.Axis)
	if err != nil {
		return nil, err
	}
	a.logger.Info("Operation Map to Table completed.")
	a.logger.Debug("The output of Operation Map to Table.", result)
	out, _ := coerce.ToArray(result)
	return out, nil
}
func convertToTable(inputMap map[string]interface{}, order []interface{}, axis int) ([][]interface{}, error) {
	//Row Order..
	if axis == 0 {
		result := make([][]interface{}, len(order))
		//row
		for index, ord := range order {
			if _, ok := inputMap[ord.(string)]; !ok {
				if givenType.String() == "string" {
					val := make([]string, ValLen)
					result[index], _ = coerce.ToArray(val)
				} else {
					val := make([]int, ValLen)
					result[index], _ = coerce.ToArray(val)
				}
			} else {
				val, _ := coerce.ToArray(inputMap[ord.(string)])
				result[index] = val
			}
		}
		return result, nil
	}
	//Column Order...
	result := make([][]interface{}, ValLen)
	for index, ord := range order {
		temp, _ := coerce.ToArray(inputMap[ord.(string)])
		//Check if the value exists
		if len(temp) != 0 {
			for key, val := range temp {
				if result[key] == nil {
					if givenType.String() == "string" {
						val := make([]string, len(order))
						result[key], _ = coerce.ToArray(val)
					} else {
						val := make([]int, len(order))
						result[key], _ = coerce.ToArray(val)
					}
				}
				result[key][index] = val
			}
		} else {
			for key, val := range result {

				if val == nil {
					//This is the  first val in column for which
					//the value in map doesn't exists
					//Hence initialize the result array with number of
					//columns
					for i := 0; i < ValLen; i++ {
						if givenType.String() == "string" {
							val := make([]string, len(order))
							result[key], _ = coerce.ToArray(val)
						} else {
							val := make([]int, len(order))
							result[key], _ = coerce.ToArray(val)
						}
					}
				}
			}
		}
	}
	return result, nil
}
