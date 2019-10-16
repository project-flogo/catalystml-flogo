package addCol2Table

import (
	"fmt"

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
	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	col := input.Col.([]interface{})
	matrix := input.Matrix.([][]interface{})

	a.logger.Info("Starting Add Column to Table operation.")
	a.logger.Debug("Inputs of operation Add Column to Table", matrix, col)

	out := matrix
	if len(col) == len(matrix) {

		for i, row := range matrix {
			out[i] = append(row, col[i])
		}
	} else {
		return out, fmt.Errorf("matrix and array of different lengths %d and %d", len(matrix), len(col))
	}

	a.logger.Info("Add Column to Table operation Completed.")
	a.logger.Debug("Output of operation Add Column To Table.", matrix, col)

	return out, nil
}
