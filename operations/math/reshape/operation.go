package reshape

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {

	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	var row, column int
	var result [][]interface{}
	in := &Input{}

	in.FromMap(inputs)
	a.logger.Info("Starting Operation reshape.")
	a.logger.Debug("The inputs for operation reshape.", in.Data)
	flatData, _ := coerce.ToArray(flattenArr(in.Data))

	if len(in.Shape) < 3 {

		row, _ = coerce.ToInt(in.Shape[0])
		column, _ = coerce.ToInt(in.Shape[1])
	}
	if row == -1 {
		return flatData, nil
	}
	if row > 0 && column <= 0 {
		rowCap := len(flatData) / row
		for i := 0; i < row; i++ {
			temp, _ := coerce.ToArray(flatData[:rowCap])
			result = append(result, temp)
			flatData = flatData[rowCap:]
		}
		return result, nil
	}

	for i := 0; i < row; i++ {
		temp, _ := coerce.ToArray(flatData[:column])
		result = append(result, temp)
		flatData = flatData[column:]
	}
	a.logger.Info("Operation reshape completed.")
	a.logger.Debug("The output for operation reshape.", result)

	return result, nil
}

func flattenArr(multiArr []interface{}) interface{} {
	var result []interface{}

	_, err := coerce.ToArray(multiArr[0])

	if err != nil {
		return multiArr
	}

	for i := 0; i < len(multiArr); i++ {

		temp, _ := coerce.ToArray(multiArr[i])

		tempResult, _ := coerce.ToArray(flattenArr(temp))

		result = append(result, tempResult...)

	}

	return result

}
