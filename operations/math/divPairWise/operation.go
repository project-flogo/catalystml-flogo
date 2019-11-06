package divPairWise

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
	//To get the inputs in the desired types.
	input := &Input{}

	input.FromMap(inputs)

	a.logger.Info("Starting operation divPairWise.")

	if inputs["matrix0"] == nil || inputs["matrix1"] == nil {

		var mtx []interface{}
		if inputs["matrix0"] != nil {
			mtx = inputs["matrix0"].([]interface{})
		} else {
			mtx = inputs["matrix1"].([]interface{})
		}

		outEdge, err := mtxZero(mtx)
		if err != nil {
			return nil, err
		}

		return outEdge, nil
	}
	mtx0 := inputs["matrix0"].([]interface{})
	mtx1 := inputs["matrix1"].([]interface{})

	out, err := mtxdivpairwise(mtx0, mtx1)
	if err != nil {
		return nil, err

	}
	a.logger.Info("Operation divPairWise completed")
	a.logger.Debug("Output of Operation divPairWise ", out)
	return out, nil
}

func mtxZero(mtx []interface{}) ([]interface{}, error) {
	var err error
	var mtxOut []interface{}

	for i := 0; i < len(mtx); i++ {
		switch v := mtx[i].(type) {
		case []interface{}:
			var tmp []interface{}
			tmp, err = mtxZero(v)
			mtxOut = append(mtxOut, tmp)
		default:
			mtxOut = append(mtxOut, 0.0)
		}

	}
	return mtxOut, err
}

func mtxdivpairwise(mtx0 []interface{}, mtx1 []interface{}) ([]interface{}, error) {
	//  Recursive function that either adds the two elements
	//   or goes to the next level of the matrices

	if len(mtx0) != len(mtx1) {
		return nil, fmt.Errorf("matrices are not the same size")
	}

	var err error
	var mtxOut []interface{}

	for i := 0; i < len(mtx0); i++ {
		switch v := mtx0[i].(type) {
		case []interface{}:
			var tmp []interface{}
			tmp, err = mtxdivpairwise(v, mtx1[i].([]interface{}))
			mtxOut = append(mtxOut, tmp)
		case int:
			temp1, _ := coerce.ToFloat32(mtx1[i])
			temp0, _ := coerce.ToFloat32(v)
			mtxOut = append(mtxOut, temp0/temp1)
		case int32:
			temp1, _ := coerce.ToFloat32(mtx1[i])
			temp0, _ := coerce.ToFloat32(mtx0[i].(int32))
			mtxOut = append(mtxOut, temp0/temp1)
		case int64:
			temp1, _ := coerce.ToFloat64(mtx1[i])
			temp0, _ := coerce.ToFloat64(mtx0[i].(int64))
			mtxOut = append(mtxOut, temp0/temp1)
		case float32:
			temp, _ := coerce.ToFloat32(mtx1[i])
			mtxOut = append(mtxOut, mtx0[i].(float32)/temp)
		case float64:
			temp, _ := coerce.ToFloat64(mtx1[i])
			mtxOut = append(mtxOut, mtx0[i].(float64)/temp)
		}
	}
	return mtxOut, err
}
