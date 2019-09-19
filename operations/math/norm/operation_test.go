package norm

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestArrayEval(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = []float32{1.0, 2.0, 2.0, 4.0, 3.0, 6.0, 4, 8.0}

	optInitConext := test.NewOperationInitContext(nil, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	assert.Nil(t, err)

	assert.Equal(t, float64(12.24744871391589), result)
}

func TestMatrixEval(t *testing.T) {

	optInitConext := test.NewOperationInitContext(nil, nil)

	inputs := make(map[string]interface{})
	inputs["data"] = [][]float32{{1.0, 2.0}, {2.0, 4.0}, {3.0, 6.0}, {4, 8.0}}

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	assert.Nil(t, err)

	assert.Equal(t, []interface{}{5.477225575051661, 10.954451150103322}, result)
}

func TestMatrixEval0(t *testing.T) {

	params := Params{Axis: 0}
	optInitConext := test.NewOperationInitContext(params, nil)

	inputs := make(map[string]interface{})
	inputs["data"] = [][]float32{{1.0, 2.0}, {2.0, 4.0}, {3.0, 6.0}, {4, 8.0}}

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	assert.Nil(t, err)

	assert.Equal(t, []interface{}{5.477225575051661, 10.954451150103322}, result)
}

func TestMatrixEval1(t *testing.T) {

	params := Params{Axis: 1}
	optInitConext := test.NewOperationInitContext(params, nil)

	inputs := make(map[string]interface{})
	inputs["data"] = [][]float32{{1.0, 2.0}, {2.0, 4.0}, {3.0, 6.0}, {4, 8.0}}

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	assert.Nil(t, err)

	assert.Equal(t, []interface{}{2.23606797749979, 4.47213595499958, 6.708203932499369, 8.94427190999916}, result)
}

func TestMatrixEval2(t *testing.T) {

	params := Params{Axis: 5}
	optInitConext := test.NewOperationInitContext(params, nil)

	inputs := make(map[string]interface{})
	inputs["data"] = [][]float32{{1.0, 2.0}, {2.0, 4.0}, {3.0, 6.0}, {4, 8.0}}

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	assert.Nil(t, err)

	assert.Equal(t, []interface{}{5.477225575051661, 10.954451150103322}, result)
}
