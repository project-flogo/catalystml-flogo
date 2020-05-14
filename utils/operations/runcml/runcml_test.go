package runcml

import (
	"testing"

	_ "github.com/project-flogo/catalystml-flogo/action"
	"github.com/project-flogo/catalystml-flogo/action/support/test"
	_ "github.com/project-flogo/operation/math"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	inputs := make(map[string]interface{})
	inputs["data"] = map[string]interface{}{"paragraph": "Abc"}
	params := map[string]interface{}{"catalystMlURI": "file://samplecml.json"}
	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)
	result, err := opt.Eval(inputs)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}
