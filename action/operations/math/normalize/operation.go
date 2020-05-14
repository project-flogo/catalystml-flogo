package normalize

import (
	"strings"

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
	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	var result interface{}

	a.logger.Info("Starting operation Normalization")
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
	a.logger.Info("Operation Normalization completed")

	return result, err

}
func calculateNorm(array []interface{}, value float32, min float32) (result []interface{}, err error) {

	//Check if the first element of the array
	//can be coerced to Float32
	_, err = coerce.ToFloat32(array[0])

	if err != nil {
		//The element present is type string.
		if strings.Contains(err.Error(), "invalid syntax") {
			return nil, nil
		}

		//If not.. consider it as an array and call itself.
		for i := 0; i < len(array); i++ {
			arr, err := coerce.ToArray(array[i])
			if err != nil {
				return nil, err
			}
			temp, err := calculateNorm(arr, value, min)
			if err != nil {
				return nil, err
			}
			result = append(result, temp)
		}
		return result, nil
	}

	temp, err := calulate1D(array, value, min)
	return temp, err

}

func calulate1D(array []interface{}, value float32, min float32) (result []interface{}, err error) {

	for key, val := range array {
		temp, err := coerce.ToFloat32(val)
		if err != nil {
			return nil, err
		}
		temp = (temp - min) / (value - min)
		array[key] = temp

	}

	return array, nil
}
