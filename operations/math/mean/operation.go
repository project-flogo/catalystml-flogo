package mean

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

func (operation *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	var result interface{}

	if !in.isFlat {
		operation.logger.Info("Matrix is...", in.Data)
		operation.logger.Info("Axis is...", operation.params.Axis)
	} else {
		operation.logger.Info("Matrix is...", in.Data.([]interface{})[0])
		operation.logger.Info("Flat array axis won't apply.")
		operation.params.Axis = -1
	}

	result, err = mean(in.Data.([]interface{}), operation.params.Axis)

	if nil != err {
		return nil, err
	}

	if in.isFlat {
		result = result.([]interface{})[0]
	}

	operation.logger.Info("Mean is..", result)

	return result, err
}

func mean(matrix []interface{}, axis int) ([]interface{}, error) {

	var result []interface{}
	var size int
	var dataAssigningIndex int

	for rowIndex, row := range matrix {
		rowArray := row.([]interface{})
		for columnIndex, column := range rowArray {
			data, _ := coerce.ToFloat64(column)
			if 1 == axis {
				if nil == result {
					size = len(rowArray)
					result = make([]interface{}, len(matrix))
				}
				dataAssigningIndex = rowIndex
			} else if 0 == axis {
				if nil == result {
					size = len(matrix)
					result = make([]interface{}, len(rowArray))
				}
				dataAssigningIndex = columnIndex
			} else {
				if nil == result {
					size = len(matrix) * len(rowArray)
					result = make([]interface{}, 1)
				}
				dataAssigningIndex = 0
			}

			if nil == result[dataAssigningIndex] {
				result[dataAssigningIndex] = data
			} else {
				result[dataAssigningIndex] = result[dataAssigningIndex].(float64) + data
			}

		}
	}

	for index, _ := range result {
		result[index] = result[index].(float64) / float64(size)
	}
	return result, nil
}
