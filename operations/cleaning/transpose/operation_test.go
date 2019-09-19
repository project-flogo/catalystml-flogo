package transpose

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestMultipleDEval(t *testing.T) {
	inputs := make(map[string]interface{})
	inputs["data"] = [][]float32{{1.0, 2.5}, {2.0, 4.1}, {3.6, 0.9}, {4, 2.3}}

	optInitConext := test.NewOperationInitContext(nil, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}
