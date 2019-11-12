package common

import (
	"fmt"

	"github.com/project-flogo/core/data/coerce"
)

type matrixOp func(val1 interface{}, val2 interface{}) (float64, error)

const (
	MatirxPairWiseAdd = iota
	MatirxPairWiseDivide
	MatirxPairWiseMulti
)

func MatrixPairWiseOperation(m0 []interface{}, m1 []interface{}, option int) ([]interface{}, error) {

	if m0 == nil || m1 == nil {

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
	var mOp matrixOp
	if option == MatirxPairWiseAdd {
		mOp = func(val1 interface{}, val2 interface{}) (float64, error) {
			temp1, err := coerce.ToFloat64(val1)
			temp2, err := coerce.ToFloat64(val2)
			return temp1 + temp2, err
		}
	} else if option == MatirxPairWiseMulti {
		mOp = func(val1 interface{}, val2 interface{}) (float64, error) {
			temp1, err := coerce.ToFloat64(val1)
			temp2, err := coerce.ToFloat64(val2)
			return temp1 * temp2, err
		}
	} else if option == MatirxPairWiseDivide {
		mOp = func(val1 interface{}, val2 interface{}) (float64, error) {
			temp1, err := coerce.ToFloat64(val1)
			temp2, err := coerce.ToFloat64(val2)
			return temp1 / temp2, err
		}
	}

	var err error
	var mtxOut []interface{}

	for i := 0; i < len(mtx0); i++ {
		switch v := mtx0[i].(type) {
		case []interface{}:
			var tmp []interface{}

			tmp, err = pairwiseOp(v, mtx1[i].([]interface{}), option)
			mtxOut = append(mtxOut, tmp)
		case int:

			val, err := mOp(v, mtx1[i])
			if err != nil {
				return nil, err
			}
			mtxOut = append(mtxOut, int(val))
		case int32:

			val, err := mOp(mtx0[i], mtx1[i])
			if err != nil {
				return nil, err
			}
			mtxOut = append(mtxOut, int32(val))
		case int64:

			val, err := mOp(mtx0[i], mtx1[i])
			if err != nil {
				return nil, err
			}
			mtxOut = append(mtxOut, int64(val))
		case float32:

			val, err := mOp(mtx0[i], mtx1[i])
			if err != nil {
				return nil, err
			}
			mtxOut = append(mtxOut, float32(val))
		case float64:

			val, err := mOp(mtx0[i], mtx1[i])
			if err != nil {
				return nil, err
			}
			mtxOut = append(mtxOut, float64(val))
		}
	}

	return mtxOut, err
}
