package norm

import (
	"math"

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

func (operation *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	var result interface{}
	operation.logger.Info("Starting Operation Norm.")
	if !in.isFlat {
		operation.logger.Debug("Matrix is...", in.Data)
		if 0 != operation.params.Axis && 1 != operation.params.Axis {
			operation.logger.Debug("Invalid axis...", operation.params.Axis, ", will set to default...0")
			operation.params.Axis = 0
		} else {
			operation.logger.Debug("Axis is...", operation.params.Axis)
		}
	} else {
		operation.logger.Debug("Matrix is...", in.Data.([]interface{})[0])
		operation.logger.Debug("Flat array axis won't apply.")
		operation.params.Axis = -1
	}

	result, err = norm(in.Data.([]interface{}), operation.params.Axis)

	if nil != err {
		return nil, err
	}

	if in.isFlat {
		result = result.([]interface{})[0]
	}

	operation.logger.Info("Operation Norm completed.")
	operation.logger.Debug("Output of Operation Norm.", result)

	return result, err
}

func norm(matrix []interface{}, axis int) ([]interface{}, error) {

	var result []interface{}
	var dataAssigningIndex int

	for rowIndex, row := range matrix {
		rowArray := row.([]interface{})
		for columnIndex, column := range rowArray {
			data, _ := coerce.ToFloat64(column)
			if 1 == axis {
				if nil == result {
					result = make([]interface{}, len(matrix))
				}
				dataAssigningIndex = rowIndex
			} else if 0 == axis {
				if nil == result {
					result = make([]interface{}, len(rowArray))
				}
				dataAssigningIndex = columnIndex
			} else {
				if nil == result {
					result = make([]interface{}, 1)
				}
				dataAssigningIndex = 0
			}

			if nil == result[dataAssigningIndex] {
				result[dataAssigningIndex] = data * data
			} else {
				result[dataAssigningIndex] = result[dataAssigningIndex].(float64) + data*data
			}

		}
	}

	for index, _ := range result {
		result[index] = math.Sqrt(result[index].(float64))
	}
	return result, nil
}
