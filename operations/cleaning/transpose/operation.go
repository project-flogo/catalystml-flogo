package transpose

import (
	"fmt"

	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = operation.Register(&Operation{}, New)
}

type Operation struct {
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	var result interface{}

	a.logger.Info("Input matrix is : ", in.Data)

	result, err = transpose(in.Data.([]interface{}))

	a.logger.Info("Transpose matrix is : ", result)

	return result, err
}

func transpose(matrix []interface{}) (result []interface{}, err error) {

	var transpose []interface{}

	for rowIndex, row := range matrix {
		//fmt.Println("Row : ", row)
		rowArray := row.([]interface{})
		if nil == transpose {
			transpose = make([]interface{}, len(rowArray))
			//fmt.Println("matrix : ", len(matrix), ", transpose : ", len(transpose))
		}

		if len(rowArray) != len(transpose) {
			return nil, fmt.Errorf("Unable to apply transpose operation - uneven column size.")
		}

		for columnIndex, column := range rowArray {
			//fmt.Println("Column : ", column)
			var newRow []interface{}
			if 0 == rowIndex {
				newRow = make([]interface{}, len(matrix))
				transpose[columnIndex] = newRow
			} else {
				newRow = transpose[columnIndex].([]interface{})
			}
			newRow[rowIndex] = column
		}
	}

	return transpose, nil
}
