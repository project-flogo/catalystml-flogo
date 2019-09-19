package mean

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

	assert.Equal(t, float64(3.75), result)
}

func TestMatrixEval(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = [][]float32{
		{1.0, 2.0},
		{2.0, 4.0},
		{3.0, 6.0},
		{4, 8.0},
	}

	optInitConext := test.NewOperationInitContext(nil, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	assert.Nil(t, err)

	assert.Equal(t, []interface{}{2.5, 5.0}, result)
}

func TestMatrixEval0(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = [][]float32{{1.0, 2.0}, {2.0, 4.0}, {3.0, 6.0}, {4, 8.0}}

	params := Params{Axis: 0}
	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	assert.Nil(t, err)

	assert.Equal(t, []interface{}{2.5, 5.0}, result)
}

func TestMatrixEval1(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = [][]float32{{1.0, 2.0}, {2.0, 4.0}, {3.0, 6.0}, {4, 8.0}}

	params := Params{Axis: 1}
	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	assert.Nil(t, err)

	assert.Equal(t, []interface{}{1.5, 3.0, 4.5, 6.0}, result)
}

func TestMatrixEval3(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = [][]float32{{1.0, 2.0}, {2.0, 4.0}, {3.0, 6.0}, {4, 8.0}}

	params := Params{Axis: 10}
	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	assert.Nil(t, err)

	assert.Equal(t, []interface{}{3.75}, result)
}
