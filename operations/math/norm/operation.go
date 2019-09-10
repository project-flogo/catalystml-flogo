package norm

import (
	"math"

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

func (this *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	var result interface{}

	if !in.isFlat {
		this.logger.Info("Matrix is...", in.Data)
		if 0 != this.params.Axis && 1 != this.params.Axis {
			this.logger.Info("Invalid axis...", this.params.Axis, ", will set to default...0")
			this.params.Axis = 0
		} else {
			this.logger.Info("Axis is...", this.params.Axis)
		}
	} else {
		this.logger.Info("Matrix is...", in.Data.([]interface{})[0])
		this.logger.Info("Flat array axis won't apply.")
		this.params.Axis = -1
	}

	result, err = norm(in.Data.([]interface{}), this.params.Axis)

	if nil != err {
		return nil, err
	}

	if in.isFlat {
		result = result.([]interface{})[0]
	}

	this.logger.Info("Norm is..", result)

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
