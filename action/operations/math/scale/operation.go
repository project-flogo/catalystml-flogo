package scale

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
	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	var result interface{}

	a.logger.Info("Starting Operation Scale.")
	a.logger.Debug("The input of Operation Scale.", in.Data)
	if _, ok := in.Data.([]interface{}); ok {
		result, err = calculateScaler(in.Data.([]interface{}), in.Scaler)
	}

	a.logger.Info("Operation Scale completed.")
	a.logger.Debug("The output of Operation Scale.", result)

	return result, err

}
func calculateScaler(array []interface{}, scaler float32) (result interface{}, err error) {

	//Check if the first element of the array
	//can be coerced to Float32
	//fmt.Println("Array....", array[0])
	_, err = coerce.ToFloat32(array[0])

	if err != nil {
		//If not.. consider it as an array and call itself.
		for i := 0; i < len(array); i++ {
			arr, _ := coerce.ToArray(array[i])
			temp, _ := calculateScaler(arr, scaler)
			array[i] = temp
		}
		return array, nil
	}
	temp := calulate1D(array, scaler)
	return temp, nil

}

func calulate1D(array []interface{}, value float32) (result []interface{}) {

	for key, val := range array {
		temp, _ := coerce.ToFloat32(val)
		array[key] = temp * value

	}

	return array
}
