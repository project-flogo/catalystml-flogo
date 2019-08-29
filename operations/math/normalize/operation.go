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

	a.logger.Info("Executing operation Normalization")
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
	a.logger.Info("Normalization finished")

	return result, err

}
func calculateNorm(array []interface{}, value float32, min float32) (result []interface{}, err error) {

	//Check if the first element of the array
	//can be coerced to Float32
	_, err = coerce.ToFloat32(array[0])

	if err != nil {
		//If not.. consider it as an array and call itself.
		for i := 0; i < len(array); i++ {
			arr, _ := coerce.ToArray(array[i])
			temp, _ := calculateNorm(arr, value, min)
			result = append(result, temp)
		}
		return result, nil
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
