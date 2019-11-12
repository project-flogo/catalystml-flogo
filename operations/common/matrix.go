package common

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
)

const (
	MatirxPairWiseAdd = iota
	MatirxPairWiseDivide
	MatirxPairWiseMulti
)

func MatrixPairWiseOperation(m0 []interface{}, m1 []interface{}, option int) ([]interface{}, error) {

	if m0 == nil || m1== nil {

		var mtx []interface{}
		if m0 != nil {

			mtx = m0
		} else {
			mtx = m1
		}

		outEdge, err := mtxZero(mtx)
		if err != nil {
			return nil, err
		}

		return outEdge, nil
	}
	mtx0 := m0
	mtx1 := m1

	out, err := pairwiseOp(mtx0, mtx1, option)
	if err != nil {
		return nil, err

	}
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

func pairwiseOp(mtx0 []interface{}, mtx1 []interface{}, option int) ([]interface{}, error) {
	//  Recursive function that either adds the two elements
	//   or goes to the next level of the matrices

	if len(mtx0) != len(mtx1) {
		return nil, fmt.Errorf("matrices are not the same size")
	}

	var err error
	var mtxOut []interface{}
	if option == MatirxPairWiseMulti {

		for i := 0; i < len(mtx0); i++ {
			switch v := mtx0[i].(type) {
			case []interface{}:
				var tmp []interface{}

				tmp, err = pairwiseOp(v, mtx1[i].([]interface{}), option)
				mtxOut = append(mtxOut, tmp)
			case int:
				temp, _ := coerce.ToInt(mtx1[i])
				mtxOut = append(mtxOut, v*temp)
			case int32:
				temp, _ := coerce.ToInt32(mtx1[i])
				mtxOut = append(mtxOut, mtx0[i].(int32)*temp)
			case int64:
				temp, _ := coerce.ToInt64(mtx1[i])
				mtxOut = append(mtxOut, mtx0[i].(int64)*temp)
			case float32:
				temp, _ := coerce.ToFloat32(mtx1[i])
				mtxOut = append(mtxOut, mtx0[i].(float32)*temp)
			case float64:
				temp, _ := coerce.ToFloat64(mtx1[i])
				mtxOut = append(mtxOut, mtx0[i].(float64)*temp)
			}
		}

	} else if option == MatirxPairWiseAdd {

		for i := 0; i < len(mtx0); i++ {
			switch v := mtx0[i].(type) {
			case []interface{}:
				var tmp []interface{}

				tmp, err = pairwiseOp(v, mtx1[i].([]interface{}), option)
				mtxOut = append(mtxOut, tmp)
			case int:
				temp, _ := coerce.ToInt(mtx1[i])
				mtxOut = append(mtxOut, v*temp)
			case int32:
				temp, _ := coerce.ToInt32(mtx1[i])
				mtxOut = append(mtxOut, mtx0[i].(int32)+temp)
			case int64:
				temp, _ := coerce.ToInt64(mtx1[i])
				mtxOut = append(mtxOut, mtx0[i].(int64)+temp)
			case float32:
				temp, _ := coerce.ToFloat32(mtx1[i])
				mtxOut = append(mtxOut, mtx0[i].(float32)+temp)
			case float64:
				temp, _ := coerce.ToFloat64(mtx1[i])
				mtxOut = append(mtxOut, mtx0[i].(float64)+temp)
			}
		}

	} else if option == MatirxPairWiseDivide {

		for i := 0; i < len(mtx0); i++ {
			switch v := mtx0[i].(type) {
			case []interface{}:
				var tmp []interface{}

				tmp, err = pairwiseOp(v, mtx1[i].([]interface{}), option)
				mtxOut = append(mtxOut, tmp)
			case int:
				temp, _ := coerce.ToInt(mtx1[i])
				mtxOut = append(mtxOut, v*temp)
			case int32:
				temp, _ := coerce.ToInt32(mtx1[i])
				mtxOut = append(mtxOut, mtx0[i].(int32)/temp)
			case int64:
				temp, _ := coerce.ToInt64(mtx1[i])
				mtxOut = append(mtxOut, mtx0[i].(int64)/temp)
			case float32:
				temp, _ := coerce.ToFloat32(mtx1[i])
				mtxOut = append(mtxOut, mtx0[i].(float32)/temp)
			case float64:
				temp, _ := coerce.ToFloat64(mtx1[i])
				mtxOut = append(mtxOut, mtx0[i].(float64)/temp)
			}
		}
	}


	return mtxOut, err
}
