package interpolateMissing

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {

	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	table["a"] = []interface{}{0.0, nil, 2.0, nil}
	table["b"] = []interface{}{nil, 2.0, 3.0, 4.0}
	table["c"] = []interface{}{-1.0, nil, nil, -4.0}
	table["d"] = []interface{}{1.0, nil, 9.0, 16.0}
	t.Log("Input of Operation Interpolate : ", table)

	inputs["data"] = table

	optInitConext := test.NewOperationInitContext(Params{}, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Output of Operation Interpolate : ", out)
}

func TestSample2(t *testing.T) {

	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	table["a"] = []interface{}{0.0, nil, 2.0, nil}
	table["b"] = []interface{}{nil, 2.0, 3.0, 4.0}
	table["c"] = []interface{}{-1.0, nil, nil, -4.0}
	table["d"] = []interface{}{1.0, nil, 9.0, 16.0}
	t.Log("Input of Operation Interpolate : ", table)

	inputs["data"] = table
	inputs["col"] = "c"

	optInitConext := test.NewOperationInitContext(Params{}, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Output of Operation Interpolate : ", out)
}

func TestSample3(t *testing.T) {

	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	table["a"] = []interface{}{0.0, nil, 2.0, nil}
	table["b"] = []interface{}{nil, 2.0, 3.0, 4.0}
	table["c"] = []interface{}{-1.0, nil, nil, -4.0}
	table["d"] = []interface{}{1.0, nil, 9.0, 16.0}
	t.Log("Input of Operation Interpolate : ", table)

	inputs["data"] = table
	inputs["col"] = "c"

	params := Params{
		How: "mean",
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Output of Operation Interpolate : ", out)
}

func TestSample4(t *testing.T) {

	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	table["a"] = []interface{}{0.0, nil, 2.0, nil}
	table["b"] = []interface{}{nil, 2.0, 3.0, 4.0}
	table["c"] = []interface{}{-1.0, nil, nil, -4.0}
	table["d"] = []interface{}{1.0, nil, 9.0, 16.0}
	t.Log("Input of Operation Interpolate : ", table)

	inputs["data"] = table
	//inputs["col"] = "c"

	params := Params{
		How:   "linear",
		Edges: "linear",
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Output of Operation Interpolate : ", out)
}

func TestSample5(t *testing.T) {

	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	table["a"] = []interface{}{0.0, nil, 2.0, nil}
	table["b"] = []interface{}{nil, 2.0, 3.0, 4.0}
	table["c"] = []interface{}{-1.0, nil, nil, -4.0}
	table["d"] = []interface{}{1.0, nil, 9.0, 16.0}
	t.Log("Input of Operation Interpolate : ", table)

	inputs["data"] = table
	inputs["col"] = "c"

	params := Params{
		How:   "linear",
		Edges: "linear",
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Output of Operation Interpolate : ", out)
}

func TestSample6(t *testing.T) {

	inputs := make(map[string]interface{})
	matrix := [][]interface{}{
		{0.0, nil, -1.0, 1.0},
		{nil, 2.0, nil, nil},
		{2.0, 3.0, nil, 9.0},
		{nil, 4.0, -4.0, 16.0},
	}

	t.Log("Input of Operation Interpolate : ", matrix)

	inputs["data"] = matrix
	inputs["col"] = "2"

	params := Params{
		How:   "linear",
		Edges: "linear",
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Output of Operation Interpolate : ", out)
}
