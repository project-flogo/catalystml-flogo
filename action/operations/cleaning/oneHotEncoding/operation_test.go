package oneHotEncoding

import (
	"fmt"
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = map[string]interface{}{"state": []interface{}{"CA", "NC", "TX"}, "sample": []interface{}{"A", "B", "C"}}

	params := Params{
		InputColumns: []interface{}{"state"},
		KeepOrig:     false,
	}

	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	fmt.Println("Result..", result)
	assert.Nil(t, err)
}

func TestSampleMultColSepFalse(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = map[string]interface{}{"state": []interface{}{"CA", "NC", "TX"}, "sample": []interface{}{"A", "B", "C"}}

	params := Params{
		InputColumns: []interface{}{"state", "sample"},
		KeepOrig:     false,
	}

	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	fmt.Println("Result..", result)
	assert.Nil(t, err)
}

func TestSampleWOutput(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = map[string]interface{}{"state": []interface{}{"CA", "NY", "TX"}, "sample": []interface{}{"A", "NY", "C"}, "NY": []interface{}{1, 2, 3}}

	params := Params{
		InputColumns:  []interface{}{"state", "sample"},
		OutputColumns: []interface{}{"CA", "NC", "TX", "NY", "A", "B", "C", "D", "E", "F"},
		KeepOrig:      true,
	}

	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	fmt.Println("Result..", result)
	assert.Nil(t, err)
}

func TestSampleMultColSepTrue(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = map[string]interface{}{"state": []interface{}{"CA", "NC", "TX"}, "sample": []interface{}{"A", "B", "C"}}

	params := Params{
		InputColumns: []interface{}{"state", "sample"},
		KeepOrig:     true,
	}

	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	fmt.Println("Result..", result)
	assert.Nil(t, err)
}

func TestSampleDefParams(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = map[string]interface{}{"state": []interface{}{"CA", "NC", "TX"}, "sample": []interface{}{"A", "B", "C"}}

	params := Params{InputColumns: []interface{}{"state"}}

	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	fmt.Println("Result..", result)
	assert.Nil(t, err)
}

func TestSampleMtx(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = []interface{}{[]interface{}{0, 2, 4}, []interface{}{6, 8, 12}, []interface{}{14, 16, 18}}

	params := Params{InputColumns: []interface{}{"0"}}

	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	fmt.Println("Result..", result)
	assert.Nil(t, err)
}

func TestSampleMtx2(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = []interface{}{[]interface{}{"a", "b", "c"}, []interface{}{"d", "e", "f"}, []interface{}{"g", "h", "i"}}

	params := Params{InputColumns: []interface{}{"0","1","2"}}

	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	fmt.Println("Result..", result)
	assert.Nil(t, err)
}
