package normalize

import (
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
	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	var result interface{}

	a.logger.Debug("Array is...", in.Data)
	if _, ok := in.Data.([]interface{}); ok {
		result, err = calculateNorm(in.Data.([]interface{}), in.Value, in.Min)
	} else {
		val, err := coerce.ToFloat32(in.Data)
		if err != nil {
			return nil, err
		}
		result = (val - in.Min) / (in.Value - in.Min)
	}

	a.logger.Debug("Norm is..", result)

	return result, err

}
func calculateNorm(array []interface{}, value float32, min float32) (result interface{}, err error) {

	_, err = coerce.ToFloat32(array[0])

	if err != nil {
		return calulate2D(array, value, min), nil
	}
	temp := calulate1D(array, value, min)
	return temp, nil

}

func calulate1D(array []interface{}, value float32, min float32) (result []interface{}) {

	for key, val := range array {
		temp, _ := coerce.ToFloat32(val)
		temp = (temp - min) / (value - min)
		array[key] = temp

	}

	return array
}

func calulate2D(array []interface{}, value float32, min float32) (result []interface{}) {

	for _, i := range array {

		val, _ := coerce.ToArray(i)
		result = append(result, calulate1D(val, value, min))
	}

	//return math.Sqrt(result)

	return result
}
