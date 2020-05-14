package replaceValue

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestMultipleDEval(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["data"] = map[string]interface{}{"state": []interface{}{"CA", "NC", "TX"}}
	inputs["replaceMap"] = map[string]interface{}{"CA": 0, "NC": 1, "TX": 2}

	params := Params{Columns: []interface{}{"state"}}

	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}
