package flatten

import (
	"errors"

	"github.com/project-flogo/cml/action/operation"
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

	multiArr := inputs["data"]

	temp, err := coerce.ToArray(multiArr)
	if err != nil {
		return nil, errors.New("Cannot Flatten non-array elements")
	}

	result := flattenArr(temp)

	if result == nil {
		return nil, nil
	}
	a.logger.Info("Flatten is..", result)

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
