package runcml

import (
	"fmt"
	"github.com/project-flogo/catalystml-flogo/action/support/test"
	_ "github.com/project-flogo/catalystml-flogo/action"
	_ "github.com/project-flogo/operation/math"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSample(t *testing.T) {
	inputs := make(map[string]interface{})
	data := make(map[string]interface{})
	data["input"] = "Abc"

	inputs["data"] = data

	params := Params{CatalystMlURI: "file://samplecml.json"}

	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)
	result, err := opt.Eval(inputs)
	fmt.Println("Result..", result)
	assert.Nil(t, err)
}