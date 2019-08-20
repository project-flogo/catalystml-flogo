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

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	var result interface{}

	a.logger.Debug("Array is...", in.Data)
	result, err = calculateNorm(in.Data.([]interface{}), 1)

	a.logger.Info("Norm is..", result)

	return result, err
}

func calculateNorm(array []interface{}, axis int) (result interface{}, err error) {

	_, err = coerce.ToFloat64(array[0])

	if err != nil {
		return calulate2D(array, axis), nil
	}
	temp := calulate1D(array)
	return []interface{}{temp}, nil

}

func calulate1D(array []interface{}) (result float64) {

	for _, val := range array {

		i, _ := coerce.ToFloat64(val)

		result = result + i*i
	}

	return math.Sqrt(result)
}

func calulate2D(array []interface{}, axis int) (result []interface{}) {

	for _, i := range array {

		val, _ := coerce.ToArray(i)
		result = append(result, calulate1D(val))
	}

	//return math.Sqrt(result)

	return result
}
