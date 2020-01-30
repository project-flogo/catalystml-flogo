package reshape

import (
	"fmt"

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
	in := &Input{}

	in.FromMap(inputs)
	a.logger.Info("Starting Operation reshape.")
	a.logger.Debug("The inputs for operation reshape. Data=", in.Data)
	a.logger.Debug("The inputs for operation reshape. Shape=", in.Shape)
	flatData, _ := coerce.ToArray(flattenArr(in.Data))

	a.logger.Info("shape length: ", len(in.Shape))

	// Checking that there aren't too many -1 values (more than 1)
	negs := 0
	ind := -1
	denom := 1
	for i, l := range in.Shape {
		lo, _ := coerce.ToInt(l)
		if lo < 0 {
			negs++
			ind = i
		} else if lo == 0 {
			return nil, fmt.Errorf("shape values of 0 are not valid")
		} else {
			denom *= lo
		}
	}
	if negs > 1 {
		return nil, fmt.Errorf("more than one shape values are -1 which is not allowed")
	}

	// Now that we know there is only one -1 . . . get the new value
	if ind >= 0 {
		newShape := float32(len(flatData)) / float32(denom)
		if newShape == float32(int32(newShape)) {
			in.Shape[ind] = newShape
		} else {
			return nil, fmt.Errorf("the value replacing -1 is not an integer")
		}
		a.logger.Info("Adusted shape array is:", in.Shape)

	}

	//building tensor see below recursive function
	tensor, err := constructTensor(flatData, in.Shape)
	if err != nil {
		return nil, err
	}
	a.logger.Info("Operation reshape completed.")
	a.logger.Debug("The output for operation reshape.", tensor)

	return tensor, nil
}

//Returns the multiple of all the integers in an array
func mul(arr []interface{}) (s int32) {
	s = 1
	for _, l := range arr {
		lo, _ := coerce.ToInt32(l)
		s *= lo
	}
	return s
}

// recursively builds array
func constructTensor(array []interface{}, shp []interface{}) (tensor []interface{}, err error) {

	arr := array
	dim, _ := coerce.ToInt32(shp[0])
	// Testing stop condition
	if len(shp) == 1 {
		return arr, nil
	}

	//self call in a loop of each dimension appended to tensor
	for i := int32(0); i < dim; i++ {
		rowCap := mul(shp[1:])
		t, err := constructTensor(arr[:rowCap], shp[1:])
		if err != nil {
			return nil, err
		}
		tensor = append(tensor, t)
		arr = arr[rowCap:]
	}
	return tensor, nil
}

//Flattens a tensor in and flattens it.
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
