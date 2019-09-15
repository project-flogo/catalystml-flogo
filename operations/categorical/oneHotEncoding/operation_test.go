package oneHotEncoding

import (
	"testing"
	"fmt"
	"github.com/project-flogo/cml/action/support/test"
	"github.com/stretchr/testify/assert"
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
