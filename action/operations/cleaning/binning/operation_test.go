package binning

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	inputs["data"] = table

	table["col1"] = []interface{}{
		30.0, 18.0, 30.0, 45.0, 31.0, 36.0, 11.0, 40.0, 53.0, 27.0,
	}

	params := Params{
		Quantile:  5,
		Labels:    []string{"1st bin", "2nd bin", "3rd bin ", "4th bin", "5th bin"},
		Column:    "col1",
		Retbins:   true,
		Precision: 2,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)
}

func TestSample2(t *testing.T) {
	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	inputs["data"] = table

	table["col1"] = []interface{}{
		30.0, 18.0, 30.0, 45.0, 31.0, 36.0, 11.0, 40.0, 53.0, 27.0,
	}

	params := Params{
		Quantile:   5,
		Column:     "col1",
		Retbins:    true,
		Precision:  2,
		Duplicates: "drop",
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)
}

func TestSample3(t *testing.T) {
	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	inputs["data"] = table

	table["col1"] = []interface{}{
		30.0, 18.0, 30.0, 45.0, 31.0, 36.0, 11.0, 40.0, 53.0, 27.0,
	}

	params := Params{
		Quantile:  5,
		Column:    "col1",
		Retbins:   false,
		Precision: 2,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)
}

func TestSample4(t *testing.T) {

	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	inputs["data"] = table

	table["col1"] = []interface{}{
		30.0, 18.0, 30.0, 45.0, 31.0, 36.0, 11.0, 40.0, 53.0, 27.0,
	}

	params := Params{
		Bins:      []float64{11, 19.4, 27.8, 36.2, 44.6, 53},
		Labels:    []string{"1st bin", "2nd bin", "3rd bin ", "4th bin", "5th bin"},
		Column:    "col1",
		Retbins:   true,
		Precision: 2,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)
}

func TestSample5(t *testing.T) {

	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	inputs["data"] = table

	table["col1"] = []interface{}{
		30.0, 18.0, 30.0, 45.0, 31.0, 36.0, 11.0, 40.0, 53.0, 27.0,
	}

	params := Params{
		Bins:       []float64{11, 19.4, 27.8, 36.2, 44.6, 53},
		Column:     "col1",
		Retbins:    true,
		Precision:  2,
		Duplicates: "drop",
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)
}

func TestSample6(t *testing.T) {

	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	inputs["data"] = table

	table["col1"] = []interface{}{
		30.0, 18.0, 30.0, 45.0, 31.0, 36.0, 11.0, 40.0, 53.0, 27.0,
	}

	params := Params{
		Bins:      []float64{11, 19.4, 27.8, 36.2, 44.6, 53},
		Labels:    []string{"1st bin", "2nd bin", "3rd bin ", "4th bin", "5th bin"},
		Column:    "col1",
		Precision: 2,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)
}
