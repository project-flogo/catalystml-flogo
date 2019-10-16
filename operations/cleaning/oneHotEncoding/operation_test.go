package oneHotEncoding

import (
	"fmt"
	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSample(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = map[string]interface{}{"state":[]interface{}{"CA","NC","TX"},"sample":"sampleSomething"}
	
	
	params := Params{Columns: []interface{}{"state"}}

	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)
	
	result, err := opt.Eval(inputs)
	fmt.Println("Result..", result)
	assert.Nil(t, err)
}
